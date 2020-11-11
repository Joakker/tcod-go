package tcod

// #include <libtcod.h>
import "C"

type Color struct {
	R, G, B byte
}

func (c *Color) getNative() C.TCOD_color_t {
	return C.TCOD_color_RGB(C.uint8_t(c.R), C.uint8_t(c.G), C.uint8_t(c.B))
}
