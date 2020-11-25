package tcod 

// #include <libtcod.h>
import "C"

type Renderer int

var (
	// Renderrenderer
	RenderSDL Renderer = C.TCOD_RENDERER_SDL
	// Renderrenderer
	RenderSDL2 Renderer = C.TCOD_RENDERER_SDL2
	// Renderrenderer
	RenderOpenGL Renderer = C.TCOD_RENDERER_OPENGL
	// Renderrenderer
	RenderOpenGL2 Renderer = C.TCOD_RENDERER_OPENGL2
)
