package tcod

/*
	#cgo pkg-config: libtcod
	#include <libtcod.h>

	void _TCOD_console_print(
				TCOD_console_t con, int x, int y, const char* s) {
			TCOD_console_printf(con, x, y, "%s", s);
		}

	void _TCOD_console_print_ex(
				TCOD_console_t con, int x, int y, TCOD_bkgnd_flag_t flag,
				TCOD_alignment_t alignment, const char* s) {
			TCOD_console_printf_ex(con, x, y, flag, alignment, "%s", s);
		}

	void _TCOD_console_print_rect(
				TCOD_console_t con, int x, int y,
				int w,int h, const char* s) {
			TCOD_console_printf_rect(
				con, x, y, w, h, "%s", s
			);
		}

	void _TCOD_console_print_frame(
				TCOD_console_t con,int x,int y,int w,int h,
				bool clear, TCOD_bkgnd_flag_t flag, const char *s) {
			TCOD_console_printf_frame(con, x, y, w, h, clear, flag, "%s", s);
		}
*/
import "C"
import "fmt"

//go:generate scripts/makeenum -v OFILE=char_const.go -v GPREFIX=Char -v DODEFAULT=true -v PACKAGE=tcod resources/chars
//go:generate scripts/makeconv resources/chars

// SetDefaultBg sets which color a console's background should have by default.
// For example:
//		c := tcod.New*Console(10, 10)
//		c.SetDefaultBg(tcod.Sky)
// makes the created *Console to use the Sky color as its background instead of
// the default black.
func (c *Console) SetDefaultBg(bg Color) {
	C.TCOD_console_set_default_background(c.console, bg.getNative())
}

// SetDefaultFg sets which color a console's foreground should have by default,
// similar to SetDefaultBg. For example:
//		c := tcod.New*Console(10, 10)
//		c.SetDefaultFg(tcod.Gold)
// makes the created console to use the Gold color as it's foreground instead of
// the default white.
func (c *Console) SetDefaultFg(fg Color) {
	C.TCOD_console_set_default_foreground(c.console, fg.getNative())
}

// Clear wipes out all cells in the console, leaving it completely clean
func (c *Console) Clear() {
	C.TCOD_console_clear(c.console)
}

// Print renders the given text to the console, using the cell at position
// (x, y) as a starting point. The console's alignment will tell whether to
// use it as the center of the text, or one of it's ends.
// You can use Go style escape sequences in the format parameter to have
// them formatted like in fmt.Sprintf. This means, however, that you are
// currently unable to format in Char* constants, as Sprintf won't render
// them as their binary values, but instead will put a space in it's place.
func (c *Console) Print(x, y int, format string, values ...interface{}) {
	final := fmt.Sprintf(format, values...)
	C._TCOD_console_print(c.console, C.int(x), C.int(y), C.CString(final))
}

// PrintEx renders the given text to the console, using the specified
// background flag and alignment.
// You can use Go style escape sequences in the format parameter to have
// them formatted like in fmt.Sprintf. This means, however, that you are
// currently unable to format in Char* constants, as Sprintf won't render
// them as their binary values, but instead will put a space in it's place.
func (c *Console) PrintEx(x, y int, flag BgFlag,
	alignment Alignment, format string, values ...interface{}) {
	final := fmt.Sprintf(format, values...)
	C._TCOD_console_print_ex(
		c.console, C.int(x), C.int(y), C.TCOD_bkgnd_flag_t(flag),
		C.TCOD_alignment_t(alignment), C.CString(final),
	)
}

// PrintRect renders a string to the console within the given bounds, that is,
// with autowrap.
// You can use Go style escape sequences in the format parameter to have
// them formatted like in fmt.Sprintf. This means, however, that you are
// currently unable to format in Char* constants, as Sprintf won't render
// them as their binary values, but instead will put a space in it's place.
func (c *Console) PrintRect(x, y, w, h int, format string, values ...interface{}) {
	final := fmt.Sprintf(format, values...)
	C._TCOD_console_print_rect(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h), C.CString(final),
	)
}

// PrintFrame prints a frame to the console. The frame has a space above
// to print the string specified by title. The argument clear tells whether
// the region within the frame should be cleared out or left alone.
// You can use Go style escape sequences in the format parameter to have
// them formatted like in fmt.Sprintf. This means, however, that you are
// currently unable to format in Char* constants, as Sprintf won't render
// them as their binary values, but instead will put a space in it's place.
func (c *Console) PrintFrame(x, y, w, h int, clear bool,
	format string, values ...interface{}) {
	final := convertString(fmt.Sprintf(format, values...))
	C._TCOD_console_print_frame(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		C.bool(clear), C.TCOD_BKGND_NONE, final,
	)
}
