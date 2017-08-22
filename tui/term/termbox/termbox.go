package termbox

import (
	"sync"

	"github.com/leeola/lab/tui/term"
	termbox "github.com/nsf/termbox-go"
)

type Term struct {
	flushLock sync.Mutex
	w, h      uint
	stacks    []*CellStack
	layers    map[uint8]*Layer
}

func New(w, h uint) (*Term, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}

	return &Term{
		w:      w,
		h:      h,
		stacks: make([]*CellStack, w*h),
		layers: map[uint8]*Layer{},
	}, nil
}

func (t *Term) Close() error {
	termbox.Close()
	return nil
}

func (t *Term) Flush() error {
	// Termbox seems to have issues with flushing concurrently, so the locking
	// prevents some broken renders and/or panics.
	t.flushLock.Lock()
	defer t.flushLock.Unlock()

	return termbox.Flush()
}

func (t *Term) Renderer(z uint8) term.Renderer {
	l, ok := t.layers[z]
	if !ok {
		l = NewLayer(t, t.w, z, t.stacks)
	}

	return l
}

type Cell struct {
	Ru rune
}

type CellStack struct {
	x, y uint
	top  uint8
	// layers is an ordered set of layer numbers, kept to allow the stack
	// to easily locate the previous layer if the top is removed.
	layers []uint8
	// TODO(leeola): benchmark/experiment with different data structures, such
	// as an array or slice, for the fastest stack storage. Map has got to be
	// slowest, right? It was chosen for early easy implementation.
	stack map[uint8]*Cell
}

func NewCellStack(x, y uint) *CellStack {
	return &CellStack{
		x:     x,
		y:     y,
		stack: map[uint8]*Cell{},
	}
}

func (cs *CellStack) Render(z uint8) error {
	// if the top is bigger then z, nothing needs
	// to be rendered, as nothing visible has changed.
	if cs.top > z {
		return nil
	}

	// TODO(leeola): we could keep track of the top cell too.
	// Eg, the rune, background and foreground. Doing so would
	// save us a lookup here.

	// get the top cell for this stack, regardless of
	// if it's the z being requested.
	c, ok := cs.stack[cs.top]
	if !ok {
		// no cell might be available in the top layer, which is not an error.
		// This happens if we remove a cell, and ask to render the top to
		// "refresh" the stack, but no other cell exists.
		//
		// In many term implementations this will result in a ghost cell rendered
		// on screen, but that's okay. This implementation lets you specify a
		// background, and that should be used if you want immediate clears.
		return nil
	}

	termbox.SetCell(int(cs.x), int(cs.y), c.Ru, termbox.ColorDefault, termbox.ColorDefault)

	return nil
}

func (cs *CellStack) SetCell(z uint8, ru rune) error {
	c := cs.stack[z]
	if c == nil {
		cs.stack[z] = &Cell{Ru: ru}
	} else {
		c.Ru = ru
	}

	if cs.top < z {
		cs.top = z
		cs.layers = append(cs.layers, z)
	} else {
		// TODO(leeola): experiment with just storing an Array of uint8 size
		// where i always equals z, and benchmark it. To compare this ordered slice
		// vs not dealing with a slice at all.
		for i, lz := range cs.layers {
			// z is already in layers slice
			if z == lz {
				break
			}

			// insert z into the list in order.
			if z < lz {
				cs.layers = append(cs.layers[:i], append([]uint8{z}, cs.layers[i:]...)...)
			}
		}
	}

	return nil
}

// NOTE(leeola): Would it make more sense to only remove the cell if it's
// the same cell instance? Eg, RemoveCell(*Cell,unt8)? This way concurrent
// assignments to the same stack don't remove eachother.. without locking,
// at least. Though, i suppose if we want true concurrency/protection,
// locking is required.
func (cs *CellStack) DelCell(z uint8) error {
	lenLayers := uint8(len(cs.layers))

	// if lenLayers is zero, there is no cell to delete.
	if lenLayers == 0 {
		return nil
	}

	// remove the cell from the stack
	delete(cs.stack, z)

	if z == cs.top {
		cs.top = lenLayers - 1

		// remove the last layer from the slice
		cs.layers = cs.layers[:cs.top]

	} else {
		// find the z layer and remove the z from the slice
		//
		// TODO(leeola): experiment with just storing an Array of uint8 size
		// where i always equals z, and benchmark it. To compare this ordered slice
		// vs not dealing with a slice at all.
		for i, lz := range cs.layers {
			if lz == z {
				// non-struct slice so this delete method is safe, i believe
				cs.layers = append(cs.layers[:i], cs.layers[i+1:]...)
				break
			}
		}
	}

	return nil
}

type Layer struct {
	flusher *Term
	w       uint
	z       uint8
	// TODO(leeola): compare different methods of keeping track of which
	// cells to were rendered.
	// Eg, hashmap vs unique slice vs non-unique slice, etc.
	rendered map[uint]bool
	stacks   []*CellStack
}

func NewLayer(flusher *Term, w uint, z uint8, stacks []*CellStack) *Layer {
	return &Layer{
		flusher:  flusher,
		w:        w,
		z:        z,
		rendered: map[uint]bool{},
		stacks:   stacks,
	}
}

func (l *Layer) Flush() error {
	for xy, rendered := range l.rendered {
		if !rendered {
			stack := l.stacks[xy]
			// remove the cell from the stack
			stack.DelCell(l.z)
			// render the stack if needed. Ie, we may have
			// removed the top cell, so we should render background cells,
			// etc.
			stack.Render(l.z)
			// remove the rendered call so it's not in the next flush cycle.
			delete(l.rendered, xy)
		} else {
			// reset each rendered value to false, for the next flush cycle.
			l.rendered[xy] = false
		}
	}

	return l.flusher.Flush()
}

func (l *Layer) Cell(x, y uint, ru rune) error {
	xy := y*l.w + x
	l.rendered[xy] = true

	stack := l.stacks[xy]
	if stack == nil {
		stack = NewCellStack(x, y)
		l.stacks[xy] = stack
	}

	if err := stack.SetCell(l.z, ru); err != nil {
		return err
	}

	return stack.Render(l.z)
}
