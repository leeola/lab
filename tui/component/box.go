package component

import (
	"github.com/leeola/lab/tui"
)

type Box struct {
	Child tui.Component
}

func (c *Box) Init(t tui.Tui) {}

func (c *Box) Render(r tui.Region) error {
	// region := Region{
	// 	Padding:    box.Padding,
	// 	PaddingTop: box.PaddingTop,
	// }

	// // TODO(leeola): replace static sizing once region has dimensions.
	width := 46
	height := 10

	d := r.Direction()
	if err := d.Cell('┌'); err != nil {
		return err
	}
	if err := d.RepeatCell('-', width-2); err != nil {
		return err
	}
	if err := d.Cell('┐'); err != nil {
		return err
	}

	d.Down()
	if err := d.RepeatCell('|', height-2); err != nil {
		return err
	}
	if err := d.Cell('┘'); err != nil {
		return err
	}

	d.Left()
	if err := d.RepeatCell('-', width-2); err != nil {
		return err
	}
	if err := d.Cell('└'); err != nil {
		return err
	}

	d.Up()
	if err := d.RepeatCell('|', height-2); err != nil {
		return err
	}

	if c.Child != nil {
		r.SubRegion(c.Child, r.Area())
	}

	return nil
}
