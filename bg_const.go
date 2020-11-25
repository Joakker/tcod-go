package tcod 

// #include <libtcod.h>
import "C"

// %s enum BgFlag
type BgFlag int

var (
	// BgNone no changes
	BgNone BgFlag = C.TCOD_BKGND_NONE
	// BgSet replaces the cell's background color with the default
	BgSet BgFlag = C.TCOD_BKGND_SET
	// BgMultiply multiplies the current background color with the default
	BgMultiply BgFlag = C.TCOD_BKGND_MULTIPLY
	// BgLighten makes the background the lightest between the old and current
	BgLighten BgFlag = C.TCOD_BKGND_LIGHTEN
	// BgDarken makes the background the darkest between the old and current
	BgDarken BgFlag = C.TCOD_BKGND_DARKEN
	// BgScreen makes the background the inverse of Multiply
	BgScreen BgFlag = C.TCOD_BKGND_SCREEN
	// BgAdd adds the current and old background colors
	BgAdd BgFlag = C.TCOD_BKGND_ADD
	// BgBurn apply a burn effect
	BgBurn BgFlag = C.TCOD_BKGND_BURN
	// BgOverlay overlays both colors
	BgOverlay BgFlag = C.TCOD_BKGND_OVERLAY
	// BgDefault default background
	BgDefault BgFlag = C.TCOD_BKGND_DEFAULT
)
