package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type Console struct {
	console C.TCOD_console_t
}

func NewConsole(w, h int32) Console {
	return Console{
		console : C.TCOD_console_new(C.int(w), C.int(h)),
	}
}

func NewConsoleASC(filename string) Console {
	return Console{
		console : C.TCOD_console_from_file(C.CString(filename)),
	}
}

func InitRoot(w, h int32, title string, fullscreen bool, renderer Renderer) *Console {
	C.TCOD_console_init_root(
		C.int(w), C.int(h), C.CString(title),
		C.bool(fullscreen), renderer.renderer,
	)
	return &Console{ console: nil }
}

func Flush() {
	C.TCOD_console_flush()
}

func Fullscreen() bool {
	return bool(C.TCOD_console_is_fullscreen())
}

func SetFullscreen(f bool) {
	C.TCOD_console_set_fullscreen(C.bool(f))
}

func SetTitle(title string) {
	C.TCOD_console_set_window_title(C.CString(title))
}

func WindowClosed() bool {
	return bool(C.TCOD_console_is_window_closed())
}

func MouseFocus() bool {
	return bool(C.TCOD_console_has_mouse_focus())
}

func WindowActive() bool {
	return bool(C.TCOD_console_is_active())
}

func CreditsScreen() {
	C.TCOD_console_credits()
}

func CreditsEmbed(x, y int32, alpha bool) bool {
	return bool(C.TCOD_console_credits_render(C.int(x), C.int(y), C.bool(alpha)))
}

func CreditsReset() {
	C.TCOD_console_credits_reset()
}

func (c *Console) LoadASC(filename string) bool {
	return bool(
		C.TCOD_console_load_asc(c.console, C.CString(filename)),
	)
}

func (c Console) SaveASC(filename string) bool {
	return bool(
		C.TCOD_console_save_asc(c.console, C.CString(filename)),
	)
}

func (c Console) Blit(x, y, w, h int32, to Console, x2, y2 int32, w2, h2 float32) {
	C.TCOD_console_blit(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		to.console, C.int(x2), C.int(y2), C.float(w2), C.float(h2),
	)
}

func (c Console) SetKeyColor(color Color) {
	C.TCOD_console_set_key_color(c.console, color.color)
}

func (c Console) Delete() {
	if c.console != nil {
		C.TCOD_console_delete(c.console)
	}
}
