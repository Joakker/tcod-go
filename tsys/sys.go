package tsys

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// SetClipboard sets the system's cliboard to the specified value
func SetClipboard(value string) {
	C.TCOD_sys_clipboard_set(C.CString(value))
}

// GetClipboard gets the clipboard contents
func GetClipboard() string {
	return C.GoString(C.TCOD_sys_clipboard_get())
}

// GetResolution returns the current screen resolution
func GetResolution() (w, h int32) {
	var cx, cy C.int
	C.TCOD_sys_get_current_resolution(&cx, &cy)
	w = int32(cx)
	h = int32(cy)
	return
}


