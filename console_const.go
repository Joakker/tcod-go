package tcod

// #include <libtcod.h>
import "C"

type Alignment int
type BgFlag int
type Renderer int

var (
	AlignCenter Alignment = C.TCOD_CENTER
	AlignLeft   Alignment = C.TCOD_LEFT
	AlignRight  Alignment = C.TCOD_RIGHT
)

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

var (
	RenderSDL     Renderer = C.TCOD_RENDERER_SDL
	RenderSDL2    Renderer = C.TCOD_RENDERER_SDL2
	RenderOpenGL  Renderer = C.TCOD_RENDERER_OPENGL
	RenderOpenGL2 Renderer = C.TCOD_RENDERER_OPENGL2
)
