package tbsp

import (
	"fmt"
)

// Bsp is a node of the Bsp tree
type Bsp struct {
	x, y, w, h int
	level      int
	l, r       *Bsp
}

// New returns a pointer to a new Bsp tree node
func New(x, y, w, h int) *Bsp {
	return &Bsp{
		x: x, y: y,
		w: w, h: h,
		level: 0,
	}
}

// GetL returns the left child of the node
func (b *Bsp) GetL() *Bsp {
	return b.l
}

// GetR returns the right child of the node
func (b *Bsp) GetR() *Bsp {
	return b.r
}

// GetX returns the position of the left edge of
// the region covered by the node
func (b *Bsp) GetX() int {
	return b.x
}

// GetY returns the position of the top edge of
// the region covered by the node
func (b *Bsp) GetY() int {
	return b.y
}

// GetW returns the width of the region covered by the node
func (b *Bsp) GetW() int {
	return b.w
}

// GetH returns the height of the region covered by the node
func (b *Bsp) GetH() int {
	return b.h
}

// Level returns how deep into the Bsp Tree the node is
func (b *Bsp) Level() int {
	return b.level
}

// Leaf returns true if the node has no children
func (b *Bsp) Leaf() bool {
	return b.l == nil && b.r == nil
}

// Contains returns true if the cell at position (x, y) is
// within the bounds of the region covered by the node
func (b *Bsp) Contains(x, y int) bool {
	return (b.x <= x && b.x+b.w < x) && (b.y <= y && b.y+b.h < y)
}

// Split partitions the region covered by the node in 2 and adds
// children that span those partitions to it. If horizontal is
// true, position must be between 0 and b.GetW(), otherwise, it
// must be between 0 and b.GetH(). If it's not within those bounds,
// Split returns an error
func (b *Bsp) Split(horizontal bool, position int) error {
	switch horizontal {
	case true:
		if position < 0 || b.w <= position {
			return fmt.Errorf("Position should be between 0 and %d", b.w)
		}
		b.l = New(b.x, b.y, position, b.h)
		b.r = New(b.x+position, b.y, b.w-position, b.h)
		b.l.level = b.level + 1
		b.r.level = b.level + 1
	case false:
		if position < 0 || b.h <= position {
			return fmt.Errorf("Position should be between 0 and %d", b.h)
		}
		b.l = New(b.x, b.y, b.w, position)
		b.r = New(b.x, b.y+position, b.w, b.h-position)
		b.l.level = b.level + 1
		b.r.level = b.level + 1
	}

	return nil
}

// Traverse performs post order traversing through the tree, starting
// from b and going down it's children to their leaves
func (b *Bsp) Traverse(callback func(node *Bsp, data *interface{}) error, data *interface{}) error {
	if !b.Leaf() {
		b.l.Traverse(callback, data)
		b.r.Traverse(callback, data)
	}
	return callback(b, data)
}
