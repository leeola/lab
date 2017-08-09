package tui

import (
	"sync"

	"github.com/leeola/errors"
	"github.com/leeola/lab/tui/term"
)

type Tui interface {
	Render()
	Quit()
}

type Region interface {
	Cell(rel_x, rel_y int, ru rune) error
	HStr(rel_x, rel_y int, s string) error
	VStr(rel_x, rel_y int, s string) error
	Direction() *Direction
	Render(Component) error
}

type Component interface {
	Init(Tui)
	Render(Region) error
}

// type Layout interface {
// 	Layout(Region) Layout
// }

type Area struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Pos struct {
	X int
	Y int
}

func Render(c Component) error {
	term, err := term.New()
	if err != nil {
		return err
	}
	defer term.Close()

	// r := &region{
	// 	area:     Area{X: 1, Y: 1, Width: 60, Height: 60},
	// }

	manager := newManager()

	tui := &tui{
		manager:   manager,
		term:      term,
		component: c,
	}

	c.Init(tui)

	// Start the render chain in another goroutine, so that we don't
	// block the manager.Wait() call
	go tui.Render()

	return manager.Wait()
}

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

// TODO(leeola): add a hierarchy setup to tui, so that each tui instance
// renders that single component.
type tui struct {
	manager   *manager
	term      term.Term
	component Component
	// children  []Component
	// children  []tui
}

func (t *tui) Render() {
	// Do nothing if we're not allowed to start.
	if allowed := t.manager.StartRender(); !allowed {
		return
	}
	defer t.manager.DoneRender()

	renderer := t.term.Renderer()
	region := &region{renderer: renderer}

	if err := t.component.Render(region); err != nil {
		// TODO(leeola): Pass the error to an error handling
		// component. This component is responsible for basically
		// just logging user errors that bubble up all the way.
		t.manager.Error(err)
		return
	}

	// Flush the rendering after each Render() request.
	if err := renderer.Flush(); err != nil {
		t.manager.Error(err)
		return
	}
}

func (r *tui) Quit() {
	r.manager.Quit()
}

type region struct {
	renderer term.Renderer
}

func (r *region) Direction() *Direction {
	return NewDirection(r)
}

func (r *region) Cell(x, y int, ru rune) error {
	r.renderer.Cell(x, y, ru)

	return nil
}

func (r *region) HStr(x, y int, s string) error {
	i := 0
	for _, ru := range s {
		if err := r.Cell(x+i, y, ru); err != nil {
			return err
		}
		i++
	}
	return nil
}

func (r *region) VStr(x, y int, s string) error {
	i := 0
	for _, ru := range s {
		if err := r.Cell(x, y+i, ru); err != nil {
			return err
		}
		i++
	}
	return nil
}

func (r *region) Render(Component) error {
	return nil
}
