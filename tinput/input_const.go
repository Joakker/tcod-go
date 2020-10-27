package tinput

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type KeyCode struct {
	code C.TCOD_keycode_t
}

var (
	KeyEscape      KeyCode = KeyCode{code: C.TCODK_ESCAPE}
	KeyBackspace   KeyCode = KeyCode{code: C.TCODK_BACKSPACE}
	KeyTab         KeyCode = KeyCode{code: C.TCODK_TAB}
	KeyEnter       KeyCode = KeyCode{code: C.TCODK_ENTER}
	KeyShift       KeyCode = KeyCode{code: C.TCODK_SHIFT}
	KeyControl     KeyCode = KeyCode{code: C.TCODK_CONTROL}
	KeyAlt         KeyCode = KeyCode{code: C.TCODK_ALT}
	KeyPause       KeyCode = KeyCode{code: C.TCODK_PAUSE}
	KeyCapslock    KeyCode = KeyCode{code: C.TCODK_CAPSLOCK}
	KeyPageup      KeyCode = KeyCode{code: C.TCODK_PAGEUP}
	KeyPagedown    KeyCode = KeyCode{code: C.TCODK_PAGEDOWN}
	KeyEnd         KeyCode = KeyCode{code: C.TCODK_END}
	KeyHome        KeyCode = KeyCode{code: C.TCODK_HOME}
	KeyUp          KeyCode = KeyCode{code: C.TCODK_UP}
	KeyLeft        KeyCode = KeyCode{code: C.TCODK_LEFT}
	KeyRight       KeyCode = KeyCode{code: C.TCODK_RIGHT}
	KeyDown        KeyCode = KeyCode{code: C.TCODK_DOWN}
	KeyPrintscreen KeyCode = KeyCode{code: C.TCODK_PRINTSCREEN}
	KeyInsert      KeyCode = KeyCode{code: C.TCODK_INSERT}
	KeyDelete      KeyCode = KeyCode{code: C.TCODK_DELETE}
	KeyLwin        KeyCode = KeyCode{code: C.TCODK_LWIN}
	KeyRwin        KeyCode = KeyCode{code: C.TCODK_RWIN}
	KeyApps        KeyCode = KeyCode{code: C.TCODK_APPS}
	KeyKpadd       KeyCode = KeyCode{code: C.TCODK_KPADD}
	KeyKpsub       KeyCode = KeyCode{code: C.TCODK_KPSUB}
	KeyKpdiv       KeyCode = KeyCode{code: C.TCODK_KPDIV}
	KeyKpmul       KeyCode = KeyCode{code: C.TCODK_KPMUL}
	KeyKpdec       KeyCode = KeyCode{code: C.TCODK_KPDEC}
	KeyKpenter     KeyCode = KeyCode{code: C.TCODK_KPENTER}
	KeyNumlock     KeyCode = KeyCode{code: C.TCODK_NUMLOCK}
	KeyScrolllock  KeyCode = KeyCode{code: C.TCODK_SCROLLLOCK}
	KeySpace       KeyCode = KeyCode{code: C.TCODK_SPACE}

	KeyF1  = KeyCode{code: C.TCODK_F1}
	KeyF2  = KeyCode{code: C.TCODK_F2}
	KeyF3  = KeyCode{code: C.TCODK_F3}
	KeyF4  = KeyCode{code: C.TCODK_F4}
	KeyF5  = KeyCode{code: C.TCODK_F5}
	KeyF6  = KeyCode{code: C.TCODK_F6}
	KeyF7  = KeyCode{code: C.TCODK_F7}
	KeyF8  = KeyCode{code: C.TCODK_F8}
	KeyF9  = KeyCode{code: C.TCODK_F9}
	KeyF10 = KeyCode{code: C.TCODK_F10}
	KeyF11 = KeyCode{code: C.TCODK_F11}
	KeyF12 = KeyCode{code: C.TCODK_F12}

	Key0 KeyCode = KeyCode{code: C.TCODK_0}
	Key1 KeyCode = KeyCode{code: C.TCODK_1}
	Key2 KeyCode = KeyCode{code: C.TCODK_2}
	Key3 KeyCode = KeyCode{code: C.TCODK_3}
	Key4 KeyCode = KeyCode{code: C.TCODK_4}
	Key5 KeyCode = KeyCode{code: C.TCODK_5}
	Key6 KeyCode = KeyCode{code: C.TCODK_6}
	Key7 KeyCode = KeyCode{code: C.TCODK_7}
	Key8 KeyCode = KeyCode{code: C.TCODK_8}
	Key9 KeyCode = KeyCode{code: C.TCODK_9}

	KeyKp0 KeyCode = KeyCode{code: C.TCODK_KP0}
	KeyKp1 KeyCode = KeyCode{code: C.TCODK_KP1}
	KeyKp2 KeyCode = KeyCode{code: C.TCODK_KP2}
	KeyKp3 KeyCode = KeyCode{code: C.TCODK_KP3}
	KeyKp4 KeyCode = KeyCode{code: C.TCODK_KP4}
	KeyKp5 KeyCode = KeyCode{code: C.TCODK_KP5}
	KeyKp6 KeyCode = KeyCode{code: C.TCODK_KP6}
	KeyKp7 KeyCode = KeyCode{code: C.TCODK_KP7}
	KeyKp8 KeyCode = KeyCode{code: C.TCODK_KP8}
	KeyKp9 KeyCode = KeyCode{code: C.TCODK_KP9}

	KeyChar KeyCode = KeyCode{code: C.TCODK_CHAR}
	KeyText KeyCode = KeyCode{code: C.TCODK_TEXT}
)
