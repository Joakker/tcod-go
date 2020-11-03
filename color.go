package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// NewColorRGB creates a new instance of a Color, with the given
// red, green and blue values. For example:
//		myGreen := tcod.NewColorRGB(0, 95, 95)
// creates a deep green.
func NewColorRGB(r, g, b uint8) Color {
	return Color{
		color: C.TCOD_color_RGB(C.uchar(r), C.uchar(g), C.uchar(b)),
	}
}

// NewColorHSV creates a new instance of a color, with the given
// hue, saturation and value. For example.
//		myGreen := tcod.NewColorHSV(180.0f, 100.0f, 37.3f)
// creates the same green as in NewColorRGB. It's recommended
// that if you don't know much about color theory, you stick
// with the RGB version, as it's easier to understand.
func NewColorHSV(h, s, v float32) Color {
	return Color{
		color: C.TCOD_color_HSV(C.float(h), C.float(s), C.float(v)),
	}
}

// GetR returns the color's Red value
func (c Color) GetR() uint8 {
	return uint8(c.color.r)
}

// GetG returns the color's Green value
func (c Color) GetG() uint8 {
	return uint8(c.color.g)
}

// GetB returns the color's Blue value
func (c Color) GetB() uint8 {
	return uint8(c.color.b)
}

// GetH returns the color's Hue
func (c Color) GetH() float32 {
	return float32(C.TCOD_color_get_hue(c.color))
}

// GetS returns the color's Saturation
func (c Color) GetS() float32 {
	return float32(C.TCOD_color_get_saturation(c.color))
}

// GetV returns the color's value
func (c Color) GetV() float32 {
	return float32(C.TCOD_color_get_value(c.color))
}

// Equal returns true if both c and o represent the same color,
// i.e., they have the same RGB values
func (c Color) Equal(o Color) bool {
	return bool(C.TCOD_color_equals(c.color, o.color))
}

// MultC multiplies the two colors, returning the result. This
// is equivalent to mixing them together. Note that their RGB
// values won't exceed 255, even if the numbers multiplied on
// their own do.
func (c Color) MultC(o Color) Color {
	return Color{
		color: C.TCOD_color_multiply(c.color, o.color),
	}
}

// MultS multiplies a color by a scalar, returning the result.
// This is equivalent to multiplying each of it's parameters
// by that scalar, making it lighter or darker. Note that the
// RGB values cannot excede 255 or go below 0.
func (c Color) MultS(s float32) Color {
	return Color{
		color: C.TCOD_color_multiply_scalar(c.color, C.float(s)),
	}
}

// Add sums two colors together, and returns the result. Similar
// to MultC, this is equivalent to mixing them together, but
// the result will be of lighter hue rather than darker.
func (c Color) Add(o Color) Color {
	return Color{
		color: C.TCOD_color_add(c.color, o.color),
	}
}

// Add substracts two colors together and returns the result. Similar
// to Add, the resulting color is darker than the operands.
func (c Color) Sub(o Color) Color {
	return Color{
		color: C.TCOD_color_subtract(c.color, o.color),
	}
}

// Lerp returns the linear interpolation between c and o, with
// coeficient coef, which is between 0 and 1. This means that
// the result is between c and o, and coef determines which one
// it is closest to.
func (c Color) Lerp(o Color, coef float32) Color {
	return Color{
		color: C.TCOD_color_lerp(c.color, o.color, C.float(coef)),
	}
}

// SetHSV defines the color given a hue, saturation and value.
func (c *Color) SetHSV(h, s, v float32) {
	C.TCOD_color_set_HSV(&c.color, C.float(h), C.float(s), C.float(v))
}
