package tcod

// #include <libtcod.h>
import "C"

// Renderer enum
type Renderer int

var (
	// RenderSDL renderer
	RenderSDL Renderer = C.TCOD_RENDERER_SDL
	// RenderSDL2 renderer
	RenderSDL2 Renderer = C.TCOD_RENDERER_SDL2
	// RenderOpenGL renderer
	RenderOpenGL Renderer = C.TCOD_RENDERER_OPENGL
	// RenderOpenGL2 renderer
	RenderOpenGL2 Renderer = C.TCOD_RENDERER_OPENGL2
)
