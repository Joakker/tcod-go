package tbsp

import (
	"fmt"
)

type Bsp struct {
	x, y, w, h int
	level      int
	l, r       *Bsp
}

func New(x, y, w, h int) *Bsp {
	return &Bsp{
		x: x, y: y,
		w: w, h: h,
		level: 0,
	}
}

func (b *Bsp) GetL() *Bsp {
	return b.l
}

func (b *Bsp) GetR() *Bsp {
	return b.r
}

func (b *Bsp) GetX() int {
	return b.x
}

func (b *Bsp) GetY() int {
	return b.y
}

func (b *Bsp) GetW() int {
	return b.w
}

func (b *Bsp) GetH() int {
	return b.h
}

func (b *Bsp) Level() int {
	return b.level
}

func (b *Bsp) Leaf() bool {
	return b.l == nil && b.r == nil
}

func (b *Bsp) Contains(x, y int) bool {
	return (b.x <= x && b.x+b.w < x) && (b.y <= y && b.y+b.h < y)
}

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

func (b *Bsp) Traverse(callback func(node *Bsp, data *interface{}) error, data *interface{}) error {
	if !b.Leaf() {
		b.l.Traverse(callback, data)
		b.r.Traverse(callback, data)
	}
	return callback(b, data)
}
