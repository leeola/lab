package tui

import (
	"sync"

	"github.com/leeola/errors"
)

type manager struct {
	closed       bool
	closedLock   sync.Mutex
	runningGroup sync.WaitGroup

	once sync.Once
	root *tui
	errC chan error
}

func newManager() *manager {
	return &manager{
		errC: make(chan error, 5),
	}
}

func (m *manager) setRoot(root *tui) {
	m.root = root
}

func (m *manager) StartRender() (allowed bool) {
	m.closedLock.Lock()
	defer m.closedLock.Unlock()
	if m.closed {
		return false
	}

	m.runningGroup.Add(1)

	return true
}

func (m *manager) DoneRender() {
	m.runningGroup.Done()
}

func (m *manager) Error(err error) {
	m.errC <- err
	m.once.Do(m.quit)
}

func (m *manager) quit() {
	// set closed to true, preventing future runners.
	//
	// This is being done outside of a lock to ensure there's never
	// a lock race between multiple render attempts and this trying to
	// close the attempt.
	//
	// A lock is still used before accessing the waitGroup, which keeps this
	// concurrently safe. See below.
	m.closed = true

	// Lock the closed variable to ensure any concurrent executions of
	// StartRender() don't start after Wait() is done.
	//
	// Eg, without lock, we could Wait() and immediately return because an
	// existing StartRender() call has not yet added it's Add(1) to the
	// WaitGroup. Locking here ensures that.
	// The closed assigning above ensures new StartRender() calls return false.
	m.closedLock.Lock()
	m.runningGroup.Wait()
	m.closedLock.Unlock()

	// TODO(leeola): add some type of unmount method to *tui, so that
	// m.Quit() triggers a full unmount chain.
	// m.root.Quit()

	close(m.errC)
}

func (m *manager) Quit() {
	m.once.Do(m.quit)
}

func (m *manager) Wait() error {
	var errs []error
	for err := range m.errC {
		errs = append(errs, err)
	}
	return errors.Join(errs)
}
