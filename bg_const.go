package tcod

// #include <libtcod.h>
import "C"

type BgFlag int

var (
	BgNone     BgFlag = C.TCOD_BKGND_NONE
	BgSet      BgFlag = C.TCOD_BKGND_SET
	BgMultiply BgFlag = C.TCOD_BKGND_MULTIPLY
	BgLighten  BgFlag = C.TCOD_BKGND_LIGHTEN
	BgDarken   BgFlag = C.TCOD_BKGND_DARKEN
	BgScreen   BgFlag = C.TCOD_BKGND_SCREEN
	BgAdd      BgFlag = C.TCOD_BKGND_ADD
	BgBurn     BgFlag = C.TCOD_BKGND_BURN
	BgOverlay  BgFlag = C.TCOD_BKGND_OVERLAY
	BgDefault  BgFlag = C.TCOD_BKGND_DEFAULT
)
