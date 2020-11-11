package tcod

// #include <libtcod.h>
import "C"

type Color struct {
	R, G, B byte
}

func (c *Color) getNative() C.TCOD_color_t {
	return C.TCOD_color_RGB(C.uchar(c.R), C.uchar(c.G), C.uchar(c.B))
}
