package component

import (
	"time"

	"github.com/leeola/lab/tui"
)

func AutoQuit(c tui.Component, durs ...time.Duration) tui.Component {
	var d time.Duration
	if len(durs) > 0 {
		d = durs[0]
	} else {
		d = 2 * time.Second
	}

	return &TimedQuit{
		Dur:   d,
		Child: c,
	}
}

type TimedQuit struct {
	tui   tui.Tui
	Dur   time.Duration
	Child tui.Component
}

func (c *TimedQuit) Init(t tui.Tui) {
	c.tui = t
}

func (c *TimedQuit) Render(r tui.Region) error {
	// Replace this with a TimerComponent or something, for now.
	go func() {
		time.Sleep(c.Dur)
		c.tui.Quit()
	}()

	// TODO(leeola): replace this direct render with r.AddComponent usage
	// once it's implement.
	return c.Child.Render(r)
}
