package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// NewColorRGB creates a Color struct with the given red, green
// and blue values
func NewColorRGB(r, g, b uint8) Color {
	return Color{
		color: C.TCOD_color_RGB(C.uchar(r), C.uchar(g), C.uchar(b)),
	}
}

// NewColorHSV creates a Color struct with the given hue, saturation
// and value
func NewColorHSV(h, s, v float32) Color {
	return Color{
		color: C.TCOD_color_HSV(C.float(h), C.float(s), C.float(v)),
	}
}

// Equal tests if two Color structs represent the same color
func (c Color) Equal(o Color) bool {
	return bool(C.TCOD_color_equals(c.color, o.color))
}

// MultC multiplies one color by another, returning the result
func (c Color) MultC(o Color) Color {
	return Color{
		color: C.TCOD_color_multiply(c.color, o.color),
	}
}

// MultS multiplies a color by a scalar s
func (c Color) MultS(s float32) Color {
	return Color{
		color: C.TCOD_color_multiply_scalar(c.color, C.float(s)),
	}
}

// Add sums two colors
func (c Color) Add(o Color) Color {
	return Color{
		color: C.TCOD_color_add(c.color, o.color),
	}
}

func (c Color) Sub(o Color) Color {
	return Color{
		color: C.TCOD_color_subtract(c.color, o.color),
	}
}

func (c Color) Lerp(o Color, coef float32) Color {
	return Color{
		color: C.TCOD_color_lerp(c.color, o.color, C.float(coef)),
	}
}

func (c *Color) SetHSV(h, s, v float32) {
	C.TCOD_color_set_HSV(&c.color, C.float(h), C.float(s), C.float(v))
}
