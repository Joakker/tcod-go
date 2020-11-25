package tcod 

// #include <libtcod.h>
import "C"

type Alignment int

var (
	// Alignmakes everything align at the center
	AlignCenter Alignment = C.TCOD_CENTER
	// Alignmakes everything align to the left
	AlignLeft Alignment = C.TCOD_LEFT
	// Alignmakes everything align to the right
	AlignRight Alignment = C.TCOD_RIGHT
)
