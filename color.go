package tcod

//go:generate scripts/makecolors -v OFILE=color_const.go resources/colors
//go:generate gofmt -w -s color_const.go

// #include <libtcod.h>
import "C"

// Color represents a colour
type Color struct {
	R, G, B byte
}

func (c *Color) getNative() C.TCOD_color_t {
	return C.TCOD_color_RGB(C.uint8_t(c.R),
		C.uint8_t(c.G), C.uint8_t(c.B))
}

// Add sums the RGB values of c and o, capping each
// one at 255
func (c Color) Add(o Color) Color {
	r := uint16(c.R) + uint16(o.R)
	g := uint16(c.G) + uint16(o.G)
	b := uint16(c.B) + uint16(o.B)

	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}

	return Color{
		R: byte(r),
		G: byte(g),
		B: byte(b),
	}
}

// Sub subtracts the RGB values of o from c, with a
// minimum of 0 for each one
func (c Color) Sub(o Color) Color {
	r := uint16(c.R) - uint16(o.R)
	g := uint16(c.G) - uint16(o.G)
	b := uint16(c.B) - uint16(o.B)

	if r > 255 {
		r = 0
	}
	if g > 255 {
		g = 0
	}
	if b > 255 {
		b = 0
	}

	return Color{
		R: byte(r),
		G: byte(g),
		B: byte(b),
	}
}
