// Package tcod provides a wrapper around the C library of the same name
package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type BgFlag = int32

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

type Renderer struct {
	renderer C.TCOD_renderer_t
}

var (
	RenderOpenGL  Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL}
	RenderOpenGL2 Renderer = Renderer{renderer: C.TCOD_RENDERER_OPENGL2}
	RenderSDL     Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL}
	RenderSDL2    Renderer = Renderer{renderer: C.TCOD_RENDERER_SDL2}
)
