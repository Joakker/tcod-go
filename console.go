package tcod

//go:generate scripts/makeenum -v PACKAGE=tcod -v OFILE=render_const.go -v CPREFIX=TCOD_RENDERER_ -v GPREFIX=Render -v GTYPE=Renderer resources/renderer
//go:generate scripts/makeenum -v PACKAGE=tcod -v OFILE=bg_const.go -v CPREFIX=TCOD_BKGND_ -v GPREFIX=Bg -v GTYPE=BgFlag resources/bgflag
//go:generate scripts/makeenum -v PACKAGE=tcod -v OFILE=align_const.go -v GPREFIX=Align -v GTYPE=Alignment resources/align

// #include <libtcod.h>
import "C"

import (
	"errors"
	"fmt"
	"os"

	"github.com/Joakker/tcod-go/assets"
	"github.com/OpenPeeDeeP/xdg"
)

const mapImage = "16x16-rogue-yun.png"

// Console is a wrapper around the C struct of the same name.
type Console struct {
	console C.TCOD_console_t
}

// NewConsole creates a console of size w by h cells
func NewConsole(w, h int) Console {
	c := Console{
		console: C.TCOD_console_new(C.int(w), C.int(h)),
	}
	c.SetDefaultFg(White)
	c.SetDefaultBg(Black)
	c.Clear()
	return c
}

// NewConsoleASC creates a console from the .asc image specified by filename
func NewConsoleASC(filename string) Console {
	c := Console{
		console: C.TCOD_console_from_file(C.CString(filename)),
	}
	c.SetDefaultFg(White)
	c.SetDefaultBg(Black)
	return c
}

// InitRoot initializes the game window, also known as the root console. It creates a console
// that is w tiles wide and h tiles high, with the given window title and using the specified
// renderer, and in fullscreen if so desired. The newly created window will make use of the
// default font image, bundled with this library in binary form, if no image has been selected
// explicitly by calling SetFontImage before calling this function.
func InitRoot(w, h int, title string, fullscreen bool, renderer Renderer) (Console, error) {
	mapPath := fmt.Sprintf("%s/%s", xdg.CacheHome(), mapImage)

	file, err := os.Create(mapPath)

	if err != nil {
		return Console{}, fmt.Errorf("Error initializing window: %w", err)
	}

	data, ok := assets.FS.String(fmt.Sprintf("/resources/%s", mapImage))

	if !ok {
		return Console{}, errors.New("Error initializing window: Cannot load image data")
	}

	file.Write([]byte(data))

	C.TCOD_console_set_custom_font(C.CString(mapPath), C.TCOD_FONT_LAYOUT_ASCII_INROW, 16, 16)
	C.TCOD_console_init_root(
		C.int(w), C.int(h), C.CString(title),
		C.bool(fullscreen), C.TCOD_renderer_t(renderer),
	)
	os.Remove(mapPath)
	c := Console{console: nil}
	c.SetDefaultFg(White)
	c.SetDefaultBg(Black)
	return c, nil
}

// Flush renders the screen and presents it to the player
func Flush() {
	C.TCOD_console_flush()
}

// Fullscreen returns true if the game window is in fullscreen mode
func Fullscreen() bool {
	return bool(C.TCOD_console_is_fullscreen())
}

// SetFullscreen makes the game window enter or exit fullscreen mode
func SetFullscreen(f bool) {
	C.TCOD_console_set_fullscreen(C.bool(f))
}

// SetTitle sets the window title to the specified argument
func SetTitle(title string) {
	C.TCOD_console_set_window_title(C.CString(title))
}

// WindowClosed returns true if the game window has been closed
func WindowClosed() bool {
	return bool(C.TCOD_console_is_window_closed())
}

// MouseFocus returns true if the mouse is within the game window
func MouseFocus() bool {
	return bool(C.TCOD_console_has_mouse_focus())
}

// WindowActive returns true if the game window is currently active
func WindowActive() bool {
	return bool(C.TCOD_console_is_active())
}

// SetFade sets the fading strength and color of the game window.
// A fade of 0 means the screen is filled, while a fade of 255
// means no fading effect. You can make a fading animation easily
// by doing something like this
//		for i := byte(0); i < 255; i-- {
//			tcod.SetFade(fade, tcod.Black)
//			tcod.Flush()
//		}
// This way, the screen will quickly fade to black. The fading
// parameters are persistent, so once you call them, they stay
// like that until you change them again.
func SetFade(fade byte, color Color) {
	C.TCOD_console_set_fade(C.uint8_t(fade), color.getNative())
}

// GetFade returns the fading strength of the game window
func GetFade() byte {
	return byte(C.TCOD_console_get_fade())
}

// GetFadeColor returns the fading color of the game window
func GetFadeColor() Color {
	c := C.TCOD_console_get_fading_color()
	return Color{R: byte(c.r), G: byte(c.g), B: byte(c.b)}
}

// CreditsScreen presents an animated screen with a message of the form "Powered
// by libtcod x.y.z". It halts the progression of the program, so it's best to
// call it as a splash screen right after the window has been created, rather than
// in the game loop. Any key press will interrupt the spash screen and give control
// back to the program.
//		root, err := tcod.InitRoot(80, 50, "The Adventures of Go", false, tcod.RenderSDL2)
//		if err != nil {
//			log.Fatal(err)
//		}
//		// The program will stop here until the animation is finished or the user
//		// presses a key
//		tcod.CreditsScreen()
//		for !tcod.WindowClosed() {
//			// Game Loop goes here
//		}
func CreditsScreen() {
	C.TCOD_console_credits()
}

// CreditsEmbed renders the message in CreditsScreen, but instead of taking the
// entire screen, it will be presented at the specified coordinates, and the
// alpha argument determines whether it should be drawn transparently, i.e., not
// overwriting existing images at that position. Similar to Print, the alignment
// of the root console will determine if the coordinates are the center of the
// image or one of it's corners. The function should be called over multiple
// frames, as every time, it renders the next frame in the animation. For example:
//		endCredits := false
//		for !tcod.WindowClosed() {
//			if !endCredits {
//				endCredits = tcod.CreditsEmbed(25, 25, true)
//			}
//		}
// The function will return false while the animation is not finished.
func CreditsEmbed(x, y int, alpha bool) bool {
	return bool(C.TCOD_console_credits_render(C.int(x), C.int(y), C.bool(alpha)))
}

// CreditsReset resets the animation generated by CreditsEmbed
func CreditsReset() {
	C.TCOD_console_credits_reset()
}

// GetW returns the width of the console
func (c *Console) GetW() int {
	if c.console == nil {
		return int(C.TCOD_console_get_width(c.console))
	}
	return int(c.console.w)
}

// GetH returns the height of the console
func (c *Console) GetH() int {
	if c.console == nil {
		return int(C.TCOD_console_get_height(c.console))
	}
	return int(c.console.h)
}

// GetDefaultBg returns the default background colour of a console
func (c Console) GetDefaultBg() Color {
	color := C.TCOD_console_get_default_background(c.console)
	return Color{R: byte(color.r), G: byte(color.g), B: byte(color.b)}
}

// GetDefaultFg returns the default foreground colour of a console
func (c Console) GetDefaultFg() Color {
	color := C.TCOD_console_get_default_foreground(c.console)
	return Color{R: byte(color.r), G: byte(color.g), B: byte(color.b)}
}

// GetCharBg returns the background color at a specific coordinate
func (c Console) GetCharBg(x, y int) Color {
	color := C.TCOD_console_get_char_background(
		c.console, C.int(x), C.int(y),
	)
	return Color{R: byte(color.r), G: byte(color.g), B: byte(color.b)}
}

// GetCharFg returns the foreground color at a specific coordinate
func (c Console) GetCharFg(x, y int) Color {
	color := C.TCOD_console_get_char_foreground(
		c.console, C.int(x), C.int(y),
	)
	return Color{R: byte(color.r), G: byte(color.g), B: byte(color.b)}
}

// GetChar returns the ascii code at a specific coordinate
func (c Console) GetChar(x, y int) byte {
	return byte(C.TCOD_console_get_char(c.console, C.int(x), C.int(y)))
}

// LoadASC loads a .asc ascii art file specified by filename, and writes the
// contents to the console, returning true if the operation was successful.
// Files with the .asc extension are created by the Ascii Paint program. For
// exaple, suppose you have a foo.asc file, so if you want to load them onto
// a console, simply call
//		c := tcod.NewConsole(40, 40)
//		c.LoadASC("foo.asc")
// and the file contents will be rendered to that console. Of course, you
// still have to blit the console onto the root one to see it.
func (c *Console) LoadASC(filename string) bool {
	return bool(
		C.TCOD_console_load_asc(c.console, C.CString(filename)),
	)
}

// SaveASC creates an Ascii Paint file from the current contents of the console,
// and saves them to the specified filename, returning true if the operation was
// successful.
func (c Console) SaveASC(filename string) bool {
	return bool(
		C.TCOD_console_save_asc(c.console, C.CString(filename)),
	)
}

// Blit renders the console's contents to the target, within the specified coordinates.
// For example:
//		c := tcod.NewConsole(100, 100)
//		c.Blit(0, 0, 10, 10, root, 5, 5, 5.0f, 5.0f)
// renders the contents of c in the coordinates from (0, 0) to (10, 10), to the root
// console starting from coordinates (5, 5), using the specified background and
// foreground alpha.
func (c Console) Blit(x, y, w, h int, target Console, x2, y2 int, bgAlpha, fgAlpha float32) {
	C.TCOD_console_blit(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		target.console, C.int(x2), C.int(y2), C.float(bgAlpha), C.float(fgAlpha),
	)
}

// SetKeyColor sets which color should be treated as transparent when blitting it
// onto another console. For example:
//		c := tcod.NewConsole(10, 10)
//		c.SetKeyColor(tcod.Magenta)
// will tell the program that it should treat any cell with Magenta background
// as transparent when blitting it to another console. The foreground remmains
// untouched.
func (c Console) SetKeyColor(color Color) {
	C.TCOD_console_set_key_color(c.console, color.getNative())
}

// Delete deletes the console from memory. In most cases you shouldn't have to
// worry about calling this method, as the Go garbage collector would take care
// of it. It only works with consoles created by the NewConsole function, not
// with the root one.
func (c Console) Delete() {
	if c.console != nil {
		C.TCOD_console_delete(c.console)
	}
}

// SetCellBg sets the background colour and flag at the given
// coordinates
func (c Console) SetCellBg(x, y int, color Color, flag BgFlag) {
	C.TCOD_console_set_char_background(c.console, C.int(x), C.int(y), color.getNative(),
		C.TCOD_bkgnd_flag_t(flag))
}

// SetCellFg sets the foreground colour at the given coordinates
func (c Console) SetCellFg(x, y int, color Color) {
	C.TCOD_console_set_char_foreground(c.console, C.int(x), C.int(y), color.getNative())
}

// SetChar renders a single character to the given coordinates
func (c Console) SetChar(x, y int, b byte) {
	C.TCOD_console_set_char(c.console, C.int(x), C.int(y), C.int(b))
}

// PutChar sets the character of a cell and it's background flag
func (c Console) PutChar(x, y int, b byte, flag BgFlag) {
	C.TCOD_console_put_char(c.console, C.int(x), C.int(y), C.int(b), C.TCOD_bkgnd_flag_t(flag))
}

// SetBgFlag sets the default BgFlag of a console
func (c Console) SetBgFlag(flag BgFlag) {
	C.TCOD_console_set_background_flag(c.console, C.TCOD_bkgnd_flag_t(flag))
}

// GetBgFlag returns the default BgFlag of a console
func (c Console) GetBgFlag() BgFlag {
	return BgFlag(C.TCOD_console_get_background_flag(c.console))
}

// SetAlign sets the console's alignment
func (c Console) SetAlign(align Alignment) {
	C.TCOD_console_set_alignment(c.console, C.TCOD_alignment_t(align))
}

// GetAlign returns the console's alignment
func (c Console) GetAlign() Alignment {
	return Alignment(C.TCOD_console_get_alignment(c.console))
}

// VLine renders a vertical line starting at (x, y), and
// extending downward for l cells
func (c Console) VLine(x, y, l int, flag BgFlag) {
	C.TCOD_console_vline(c.console, C.int(x), C.int(y), C.int(l), C.TCOD_bkgnd_flag_t(flag))
}

// HLine renders a horizontal line starting at (x, y) and
// extending leftward for l cells
func (c Console) HLine(x, y, l int, flag BgFlag) {
	C.TCOD_console_hline(c.console, C.int(x), C.int(y), C.int(l), C.TCOD_bkgnd_flag_t(flag))
}
