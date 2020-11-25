package tcod 

// #include <libtcod.h>
import "C"

// %s enum Alignment
type Alignment int

var (
	// AlignCenter makes everything align at the center
	AlignCenter Alignment = C.TCOD_CENTER
	// AlignLeft makes everything align to the left
	AlignLeft Alignment = C.TCOD_LEFT
	// AlignRight makes everything align to the right
	AlignRight Alignment = C.TCOD_RIGHT
)
