package tsys

// #cgo LDFLAGS: -ltcod -lSDL2
// #include <SDL2/SDL.h>
// #include <libtcod.h>
import "C"
import "github.com/Joakker/tcod-go"

// SetClipboard sets the system's cliboard to the specified value
func SetClipboard(value string) {
	C.SDL_SetClipboardText(C.CString(value))
}

// GetClipboard gets the clipboard contents
func GetClipboard() string {
	return C.GoString(C.SDL_GetClipboardText())
}

// GetResolution returns the current screen resolution
func GetResolution() (w, h int32) {
	var cx, cy C.int
	C.TCOD_sys_get_current_resolution(&cx, &cy)
	w = int32(cx)
	h = int32(cy)
	return
}

// SetRenderer dynamically changes the renderer used by the library
func SetRenderer(renderer tcod.Renderer) {
	C.TCOD_sys_set_renderer(convertRenderer(renderer))
}

/*
	SaveScreenshot saves a PNG or BMP file with the current game screen.
	The format depends on the file extension that the path ends with.
*/
func SaveScreenshot(filename string) {
	C.TCOD_sys_save_screenshot(C.CString(filename))
}

// convertRenderer converts from tcod.Renderer to the C representation
func convertRenderer(renderer tcod.Renderer) C.TCOD_renderer_t {
	switch renderer {
	case tcod.RenderOpenGL:
		return C.TCOD_RENDERER_OPENGL
	case tcod.RenderOpenGL2:
		return C.TCOD_RENDERER_OPENGL2
	case tcod.RenderSDL:
		return C.TCOD_RENDERER_SDL
	case tcod.RenderSDL2:
		return C.TCOD_RENDERER_SDL2
	default:
		return C.TCOD_RENDERER_SDL2
	}
}
