package main

import (
	"fmt"
	"time"

	"github.com/leeola/lab/tui"
	. "github.com/leeola/lab/tui/component"
)

type CustomComponent struct {
	secondsTilClose int
}

func (c *CustomComponent) Init(t tui.Tui) {
	c.secondsTilClose = 15

	for ; c.secondsTilClose > 0; c.secondsTilClose-- {
		t.Render()
		time.Sleep(time.Millisecond * 300)
	}

	t.Quit()
}

func (c *CustomComponent) Render(r tui.Region) error {
	if err := r.Fill('~'); err != nil {
		return err
	}

	if c.secondsTilClose > 12 {
		return nil
	}

	return r.SubRegion(
		Text(fmt.Sprintf("Seconds Till Quit: %d", c.secondsTilClose)),
		tui.Area{
			Z: 2,
		},
	)
}

func main() {
	if err := tui.Render(&CustomComponent{}); err != nil {
		panic(err)
	}
}
