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

func (c Console) SetDefaultBg(bg Color) {
	C.TCOD_console_set_default_background(c.console, bg.color)
}

func (c Console) SetDefaultFg(fg Color) {
	C.TCOD_console_set_default_foreground(c.console, fg.color)
}

func (c Console) Clear() {
	C.TCOD_console_clear(c.console)
}

// func (c Console) SetCellBg(x, y int32, bg Color, flag BgFlag) {
// 	C.TCOD_console_set_char_background(
// 		c.console, C.int(x), C.int(y), bg.color, flag,
// 	)
// }

// func (c Console) SetCellFg(x, y int32, fg Color, flag BgFlag) {
// 	C.TCOD_console_set_char_foreground(
// 		c.console, C.int(x), C.int(y), fg.color, flag,
// 	)
// }

// func (c Console) SetCell(x, y int32, char byte) {
// 	C.TCOD_console_set_char(
// 		c.console, C.int(x), C.int(y), C.char(char),
// 	)
// }

func (c Console) PrintFrame(x, y, w, h int32, clear bool, title string) {
	C._TCOD_console_print_frame(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		C.bool(clear), C.TCOD_BKGND_NONE,C.CString(title),
	)
}

func (c Console) Print(x, y int32, text string) {
	C._TCOD_console_print(
		c.console, C.int(x), C.int(y), C.CString(text),
	)
}
