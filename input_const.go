package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type KeyCode = int32

var (
	KeyEscape      KeyCode = C.TCODK_ESCAPE
	KeyBackspace   KeyCode = C.TCODK_BACKSPACE
	KeyTab         KeyCode = C.TCODK_TAB
	KeyEnter       KeyCode = C.TCODK_ENTER
	KeyShift       KeyCode = C.TCODK_SHIFT
	KeyControl     KeyCode = C.TCODK_CONTROL
	KeyAlt         KeyCode = C.TCODK_ALT
	KeyPause       KeyCode = C.TCODK_PAUSE
	KeyCapslock    KeyCode = C.TCODK_CAPSLOCK
	KeyPageup      KeyCode = C.TCODK_PAGEUP
	KeyPagedown    KeyCode = C.TCODK_PAGEDOWN
	KeyEnd         KeyCode = C.TCODK_END
	KeyHome        KeyCode = C.TCODK_HOME
	KeyUp          KeyCode = C.TCODK_UP
	KeyLeft        KeyCode = C.TCODK_LEFT
	KeyRight       KeyCode = C.TCODK_RIGHT
	KeyDown        KeyCode = C.TCODK_DOWN
	KeyPrintscreen KeyCode = C.TCODK_PRINTSCREEN
	KeyInsert      KeyCode = C.TCODK_INSERT
	KeyDelete      KeyCode = C.TCODK_DELETE
	KeyLwin        KeyCode = C.TCODK_LWIN
	KeyRwin        KeyCode = C.TCODK_RWIN
	KeyApps        KeyCode = C.TCODK_APPS
	KeyKpadd       KeyCode = C.TCODK_KPADD
	KeyKpsub       KeyCode = C.TCODK_KPSUB
	KeyKpdiv       KeyCode = C.TCODK_KPDIV
	KeyKpmul       KeyCode = C.TCODK_KPMUL
	KeyKpdec       KeyCode = C.TCODK_KPDEC
	KeyKpenter     KeyCode = C.TCODK_KPENTER
	KeyNumlock     KeyCode = C.TCODK_NUMLOCK
	KeyScrolllock  KeyCode = C.TCODK_SCROLLLOCK
	KeySpace       KeyCode = C.TCODK_SPACE

	KeyF1  KeyCode = C.TCODK_F1
	KeyF2  KeyCode = C.TCODK_F2
	KeyF3  KeyCode = C.TCODK_F3
	KeyF4  KeyCode = C.TCODK_F4
	KeyF5  KeyCode = C.TCODK_F5
	KeyF6  KeyCode = C.TCODK_F6
	KeyF7  KeyCode = C.TCODK_F7
	KeyF8  KeyCode = C.TCODK_F8
	KeyF9  KeyCode = C.TCODK_F9
	KeyF10 KeyCode = C.TCODK_F10
	KeyF11 KeyCode = C.TCODK_F11
	KeyF12 KeyCode = C.TCODK_F12

	Key0 KeyCode = C.TCODK_0
	Key1 KeyCode = C.TCODK_1
	Key2 KeyCode = C.TCODK_2
	Key3 KeyCode = C.TCODK_3
	Key4 KeyCode = C.TCODK_4
	Key5 KeyCode = C.TCODK_5
	Key6 KeyCode = C.TCODK_6
	Key7 KeyCode = C.TCODK_7
	Key8 KeyCode = C.TCODK_8
	Key9 KeyCode = C.TCODK_9

	KeyKp0 KeyCode = C.TCODK_KP0
	KeyKp1 KeyCode = C.TCODK_KP1
	KeyKp2 KeyCode = C.TCODK_KP2
	KeyKp3 KeyCode = C.TCODK_KP3
	KeyKp4 KeyCode = C.TCODK_KP4
	KeyKp5 KeyCode = C.TCODK_KP5
	KeyKp6 KeyCode = C.TCODK_KP6
	KeyKp7 KeyCode = C.TCODK_KP7
	KeyKp8 KeyCode = C.TCODK_KP8
	KeyKp9 KeyCode = C.TCODK_KP9

	KeyChar KeyCode = C.TCODK_CHAR
)
