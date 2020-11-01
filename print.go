package tcod

/*
	#cgo LDFLAGS: -ltcod
	#include <libtcod.h>

	void _TCOD_console_print_frame(
				TCOD_console_t con,int x,int y,int w,int h,
				bool clear, TCOD_bkgnd_flag_t flag, const char *s) {
			TCOD_console_printf_frame(con, x, y, w, h, clear, flag, "%s", s);
		}

	void _TCOD_console_print(
				TCOD_console_t con, int x, int y, const char* s) {
			TCOD_console_printf(con, x, y, "%s", s);
		}
*/
import "C"

/*
	SetDefaultBg sets which color a console's background should have by default.
	For example:

		c := tcod.NewConsole(10, 10)
		c.SetDefaultBg(tcod.Sky)

	makes the created Console to use the Sky color as its background instead of
	the default black.
*/
func (c Console) SetDefaultBg(bg Color) {
	C.TCOD_console_set_default_background(c.console, bg.color)
}

/*
	SetDefaultFg sets which color a console's foreground should have by default,
	similar to SetDefaultBg. For example:

		c := tcod.NewConsole(10, 10)
		c.SetDefaultFg(tcod.Gold)

	makes the created console to use the Gold color as it's foreground instead of
	the default white.
*/
func (c Console) SetDefaultFg(fg Color) {
	C.TCOD_console_set_default_foreground(c.console, fg.color)
}

// Clear wipes out all cells in the console, leaving it completely clean
func (c Console) Clear() {
	C.TCOD_console_clear(c.console)
}

/*
	PrintFrame prints a frame to the console. The frame has a space above
	to print the string specified by title. The argument clear tells whether
	the region within the frame should be cleared out or left alone.
*/
func (c Console) PrintFrame(x, y, w, h int, clear bool, title string) {
	C._TCOD_console_print_frame(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		C.bool(clear), C.TCOD_BKGND_NONE, C.CString(title),
	)
}

/*
	Print renders the given text to the console, using the cell at position
	(x, y) as a starting point. The console's alignment will tell whether to
	use it as the center of the text, or one of it's ends.
*/
func (c Console) Print(x, y int, text string) {
	C._TCOD_console_print(c.console, C.int(x), C.int(y), C.CString(text))
}
