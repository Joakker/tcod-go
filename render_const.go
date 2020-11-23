package tcod

// #include <libtcod.h>
import "C"

type Renderer int

var (
	RenderSDL     Renderer = C.TCOD_RENDERER_SDL
	RenderSDL2    Renderer = C.TCOD_RENDERER_SDL2
	RenderOpenGL  Renderer = C.TCOD_RENDERER_OPENGL
	RenderOpenGL2 Renderer = C.TCOD_RENDERER_OPENGL2
)
