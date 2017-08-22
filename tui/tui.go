package tui

import (
	"github.com/leeola/lab/tui/term"
	"github.com/leeola/lab/tui/term/termbox"
)

type Tui interface {
	Render()
	Quit()
}

type Region interface {
	Area() Area
	Cell(rel_x, rel_y int, ru rune) error
	HStr(rel_x, rel_y int, s string) error
	VStr(rel_x, rel_y int, s string) error
	Direction() *Direction
	SubRegion(c Component, a Area) error
	Fill(rune) error
}

type Component interface {
	Init(Tui)
	Render(r Region) error
}

// type Layout interface {
// 	Layout(Region) Layout
// }

type Area struct {
	Z      int
	RelX   int
	RelY   int
	Width  int
	Height int
}

type Pos struct {
	X int
	Y int
}

func Render(c Component) error {
	term, err := termbox.New(60, 30)
	if err != nil {
		return err
	}
	defer term.Close()

	manager := newManager()

	tui := &tui{
		manager:   manager,
		term:      term,
		component: c,
		area:      Area{RelX: 1, RelY: 1, Width: 60, Height: 60},
	}

	go func() {
		c.Init(tui)

		// Start the render chain in another goroutine, so that we don't
		// block the manager.Wait() call
		tui.Render()
	}()

	return manager.Wait()
}

// TODO(leeola): add a hierarchy setup to tui, so that each tui instance
// renders that single component.
type tui struct {
	manager   *manager
	term      term.Term
	component Component
	area      Area
	// children  []Component
	// children  []tui
}

func (t *tui) Render() {
	// Do nothing if we're not allowed to start.
	if allowed := t.manager.StartRender(); !allowed {
		return
	}
	defer t.manager.DoneRender()

	region := newRegion(t.term, t.area)

	if err := t.component.Render(region); err != nil {
		// TODO(leeola): Pass the error to an error handling
		// component. This component is responsible for basically
		// just logging user errors that bubble up all the way.
		t.manager.Error(err)
		return
	}

	// Flush the rendering after each Render() request.
	if err := region.Flush(); err != nil {
		t.manager.Error(err)
		return
	}
}

func (r *tui) Quit() {
	r.manager.Quit()
}

type region struct {
	area     Area
	term     term.Term
	renderer term.Renderer
	children map[Component]struct{}
}

func newRegion(t term.Term, a Area) *region {
	return &region{
		area:     a,
		term:     t,
		renderer: t.Renderer(uint8(a.Z)),
		children: map[Component]struct{}{},
	}
}

func (r *region) Area() Area {
	return r.area
}

func (r *region) SubRegion(c Component, a Area) error {
	// TODO(leeola): Call Init() if it's not yet been in the map.
	// Which means we need to track pre-existing children in this map
	// as well.

	r.children[c] = struct{}{}

	return c.Render(newRegion(r.term, a))
}

func (r *region) Children() map[Component]struct{} {
	return r.children
}

func (r *region) Direction() *Direction {
	return NewDirection(r)
}

func (r *region) Cell(x, y int, ru rune) error {
	r.renderer.Cell(uint(x), uint(y), ru)

	return nil
}

func (r *region) Flush() error {
	return r.renderer.Flush()
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

func (r *region) Fill(ru rune) error {
	// TODO(leeola): make the fill relative
	// for x := 0; x < r.area.Width; x++ {
	// 	for y := 0; y < r.area.Height; y++ {
	for x := 0; x < 30; x++ {
		for y := 0; y < 10; y++ {
			if err := r.renderer.Cell(uint(x), uint(y), ru); err != nil {
				return err
			}
		}
	}

	return nil
}
