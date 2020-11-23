package tcod

// #include <libtcod.h>
import "C"

type Alignment int

var (
	AlignCenter Alignment = C.TCOD_CENTER
	AlignLeft   Alignment = C.TCOD_LEFT
	AlignRight  Alignment = C.TCOD_RIGHT
)
