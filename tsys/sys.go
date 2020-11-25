package tsys

// #cgo pkg-config: libtcod sdl2
// #include <SDL.h>
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
	w, h = int32(cx), int32(cy)
	return
}

// ForceFullscreenResolution forces the window to take the
// resolution w x h when in fullscreen mode
func ForceFullscreenResolution(w, h int32) {
	C.TCOD_sys_force_fullscreen_resolution(C.int(w), C.int(h))
}

// GetFullscreenOffset returns the offset of the window
// in fullscreen mode.
func GetFullscreenOffset() (x, y int32) {
	var cx, cy C.int
	C.TCOD_sys_get_fullscreen_offsets(&cx, &cy)
	x, y = int32(cx), int32(cy)
	return
}

// GetCharSize returns the width and height of a single cell
// in pixels
func GetCharSize() (w, h int32) {
	var cx, cy C.int
	C.TCOD_sys_get_char_size(&cx, &cy)
	w, h = int32(cx), int32(cy)
	return
}

// SetRenderer dynamically changes the renderer used by the library
func SetRenderer(renderer tcod.Renderer) {
	C.TCOD_sys_set_renderer(C.TCOD_renderer_t(renderer))
}

// SaveScreenshot saves a PNG or BMP file with the current game screen.
// The format depends on the file extension that the path ends with.
func SaveScreenshot(filename string) {
	C.TCOD_sys_save_screenshot(C.CString(filename))
}
