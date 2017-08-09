package tui

const (
	right int = iota
	down
	left
	up
)

type Direction struct {
	x, y int
	dir  int
	r    Region
}

func NewDirection(r Region) *Direction {
	return &Direction{
		r: r,
	}
}

func (d *Direction) forward() {
	switch d.dir {
	case right:
		d.x++
	case down:
		d.y++
	case left:
		d.x--
	case up:
		d.y--
	}
}

func (d *Direction) backward() {
	switch d.dir {
	case right:
		d.x--
	case down:
		d.y--
	case left:
		d.x++
	case up:
		d.y++
	}
}

func (d *Direction) Cell(ru rune) error {
	err := d.r.Cell(d.x, d.y, ru)

	switch d.dir {
	case right:
		d.x++
	case down:
		d.y++
	case left:
		d.x--
	case up:
		d.y--
	}

	return err
}

func (d *Direction) RepeatCell(ru rune, n int) error {
	for i := 0; i < n; i++ {
		if err := d.Cell(ru); err != nil {
			return err
		}
	}
	return nil
}

func (d *Direction) Up() {
	d.backward()
	d.dir = up
	d.forward()
}

func (d *Direction) Down() {
	d.backward()
	d.dir = down
	d.forward()
}

func (d *Direction) Right() {
	d.backward()
	d.dir = right
	d.forward()
}

func (d *Direction) Left() {
	d.backward()
	d.dir = left
	d.forward()
}
