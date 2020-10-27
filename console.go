// Package tcod provides a wrapper around the C library of the same name
package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

import (
	"errors"
	"fmt"
	"os"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/Joakker/tcod-go/assets"
)

const mapImage = "terminal.png"

// Console is a wrapper around the C struct of the same name.
type Console struct {
	console C.TCOD_console_t
}

// NewConsole creates a console of size w by h cells
func NewConsole(w, h int32) Console {
	return Console{
		console : C.TCOD_console_new(C.int(w), C.int(h)),
	}
}

// NewConsoleASC creates a console from the .asc image specified by filename
func NewConsoleASC(filename string) Console {
	return Console{
		console : C.TCOD_console_from_file(C.CString(filename)),
	}
}

// InitRoot initializes the game window, including the font image
func InitRoot(w, h int32, title string, fullscreen bool, renderer Renderer) (*Console, error) {
	mapPath := fmt.Sprintf("%s/%s", xdg.CacheHome(), mapImage)
	file, err := os.Create(mapPath)
	if err != nil {
		return nil, fmt.Errorf("Error initializing window: %w\n", err)
	}
	data, ok := assets.FS.String(fmt.Sprintf("/%s", mapImage))
	if !ok {
		return nil, errors.New("Error initializing window: Cannot load image data")
	}
	file.Write([]byte(data))
	C.TCOD_console_set_custom_font(C.CString(mapPath), C.TCOD_FONT_LAYOUT_ASCII_INCOL, 16, 16)
	C.TCOD_console_init_root(
		C.int(w), C.int(h), C.CString(title),
		C.bool(fullscreen), renderer.renderer,
	)
	os.Remove(mapPath)
	return &Console{ console: nil }, nil
}

// Flush renders the screen and presents it to the player
func Flush() {
	C.TCOD_console_flush()
}

// Fullscreen returns true if the game window is in fullscreen mode
func Fullscreen() bool {
	return bool(C.TCOD_console_is_fullscreen())
}

// Fullscreen makes the game window enter or exit fullscreen mode
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

// CreditsScreen renders a screen with a version message
func CreditsScreen() {
	C.TCOD_console_credits()
}

// CreditsEmbed renders a version message to the (x, y) coordinates in the root console
func CreditsEmbed(x, y int32, alpha bool) bool {
	return bool(C.TCOD_console_credits_render(C.int(x), C.int(y), C.bool(alpha)))
}

// CreditsReset resets the animation generated by CreditsEmbed
func CreditsReset() {
	C.TCOD_console_credits_reset()
}

// LoadASC loads a .asc file's contents to the console
func (c *Console) LoadASC(filename string) bool {
	return bool(
		C.TCOD_console_load_asc(c.console, C.CString(filename)),
	)
}

// SaveASC saves a .asc file specified by filename, with the console's contents
func (c Console) SaveASC(filename string) bool {
	return bool(
		C.TCOD_console_save_asc(c.console, C.CString(filename)),
	)
}

// Blit renders the console's contents in the specified coordinates
// to the target console's coordinates
func (c Console) Blit(x, y, w, h int32, to Console, x2, y2 int32, w2, h2 float32) {
	C.TCOD_console_blit(
		c.console, C.int(x), C.int(y), C.int(w), C.int(h),
		to.console, C.int(x2), C.int(y2), C.float(w2), C.float(h2),
	)
}

// SetKeyColor sets the transparent color for the console
func (c Console) SetKeyColor(color Color) {
	C.TCOD_console_set_key_color(c.console, color.color)
}

// Delete deletes the console from memory. Doesn't work on the root console
func (c Console) Delete() {
	if c.console != nil {
		C.TCOD_console_delete(c.console)
	}
}
