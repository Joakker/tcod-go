package tcod 

// #include <libtcod.h>
import "C"

type BgFlag int

var (
	// Bgno changes
	BgNone BgFlag = C.TCOD_BKGND_NONE
	// Bgreplaces the cell's background color with the default
	BgSet BgFlag = C.TCOD_BKGND_SET
	// Bgmultiplies the current background color with the default
	BgMultiply BgFlag = C.TCOD_BKGND_MULTIPLY
	// Bgmakes the background the lightest between the old and current
	BgLighten BgFlag = C.TCOD_BKGND_LIGHTEN
	// Bgmakes the background the darkest between the old and current
	BgDarken BgFlag = C.TCOD_BKGND_DARKEN
	// Bgmakes the background the inverse of Multiply
	BgScreen BgFlag = C.TCOD_BKGND_SCREEN
	// Bgadds the current and old background colors
	BgAdd BgFlag = C.TCOD_BKGND_ADD
	// Bgapply a burn effect
	BgBurn BgFlag = C.TCOD_BKGND_BURN
	// Bgoverlays both colors
	BgOverlay BgFlag = C.TCOD_BKGND_OVERLAY
	// Bgdefault background
	BgDefault BgFlag = C.TCOD_BKGND_DEFAULT
)
