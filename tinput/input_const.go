package tinput

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

// KeyCode represents the key that the user is triggering
type KeyCode struct {
	code C.TCOD_keycode_t
}

var (

	// Key escape
	KeyEscape KeyCode = KeyCode{code: C.TCODK_ESCAPE}
	// Key backspace
	KeyBackspace KeyCode = KeyCode{code: C.TCODK_BACKSPACE}
	// Key tab
	KeyTab KeyCode = KeyCode{code: C.TCODK_TAB}
	// Key enter
	KeyEnter KeyCode = KeyCode{code: C.TCODK_ENTER}
	// Key shift
	KeyShift KeyCode = KeyCode{code: C.TCODK_SHIFT}
	// Key control
	KeyControl KeyCode = KeyCode{code: C.TCODK_CONTROL}
	// Key alt
	KeyAlt KeyCode = KeyCode{code: C.TCODK_ALT}
	// Key pause
	KeyPause KeyCode = KeyCode{code: C.TCODK_PAUSE}
	// Key capslock
	KeyCapslock KeyCode = KeyCode{code: C.TCODK_CAPSLOCK}
	// Key pageup
	KeyPageup KeyCode = KeyCode{code: C.TCODK_PAGEUP}
	// Key pagedown
	KeyPagedown KeyCode = KeyCode{code: C.TCODK_PAGEDOWN}
	// Key end
	KeyEnd KeyCode = KeyCode{code: C.TCODK_END}
	// Key home
	KeyHome KeyCode = KeyCode{code: C.TCODK_HOME}
	// Key up
	KeyUp KeyCode = KeyCode{code: C.TCODK_UP}
	// Key left
	KeyLeft KeyCode = KeyCode{code: C.TCODK_LEFT}
	// Key right
	KeyRight KeyCode = KeyCode{code: C.TCODK_RIGHT}
	// Key down
	KeyDown KeyCode = KeyCode{code: C.TCODK_DOWN}
	// Key printscreen
	KeyPrintscreen KeyCode = KeyCode{code: C.TCODK_PRINTSCREEN}
	// Key insert
	KeyInsert KeyCode = KeyCode{code: C.TCODK_INSERT}
	// Key delete
	KeyDelete KeyCode = KeyCode{code: C.TCODK_DELETE}
	// Key lwin
	KeyLwin KeyCode = KeyCode{code: C.TCODK_LWIN}
	// Key rwin
	KeyRwin KeyCode = KeyCode{code: C.TCODK_RWIN}
	// Key apps
	KeyApps KeyCode = KeyCode{code: C.TCODK_APPS}
	// Key kpAdd
	KeyKpAdd KeyCode = KeyCode{code: C.TCODK_KPADD}
	// Key kpSub
	KeyKpSub KeyCode = KeyCode{code: C.TCODK_KPSUB}
	// Key kpDiv
	KeyKpDiv KeyCode = KeyCode{code: C.TCODK_KPDIV}
	// Key kpMul
	KeyKpMul KeyCode = KeyCode{code: C.TCODK_KPMUL}
	// Key kpDec
	KeyKpDec KeyCode = KeyCode{code: C.TCODK_KPDEC}
	// Key kpEnter
	KeyKpEnter KeyCode = KeyCode{code: C.TCODK_KPENTER}
	// Key numlock
	KeyNumlock KeyCode = KeyCode{code: C.TCODK_NUMLOCK}
	// Key scrolllock
	KeyScrolllock KeyCode = KeyCode{code: C.TCODK_SCROLLLOCK}
	// Key space
	KeySpace KeyCode = KeyCode{code: C.TCODK_SPACE}
	// Key text
	KeyText KeyCode = KeyCode{code: C.TCODK_TEXT}
	// Key char
	KeyChar KeyCode = KeyCode{code: C.TCODK_CHAR}
	// Key Function 1
	KeyF1 KeyCode = KeyCode{code: C.TCODK_F1}
	// Key Function 2
	KeyF2 KeyCode = KeyCode{code: C.TCODK_F2}
	// Key Function 3
	KeyF3 KeyCode = KeyCode{code: C.TCODK_F3}
	// Key Function 4
	KeyF4 KeyCode = KeyCode{code: C.TCODK_F4}
	// Key Function 5
	KeyF5 KeyCode = KeyCode{code: C.TCODK_F5}
	// Key Function 6
	KeyF6 KeyCode = KeyCode{code: C.TCODK_F6}
	// Key Function 7
	KeyF7 KeyCode = KeyCode{code: C.TCODK_F7}
	// Key Function 8
	KeyF8 KeyCode = KeyCode{code: C.TCODK_F8}
	// Key Function 9
	KeyF9 KeyCode = KeyCode{code: C.TCODK_F9}
	// Key Function 10
	KeyF10 KeyCode = KeyCode{code: C.TCODK_F10}
	// Key Function 11
	KeyF11 KeyCode = KeyCode{code: C.TCODK_F11}
	// Key Function 12
	KeyF12 KeyCode = KeyCode{code: C.TCODK_F12}
	// Key 0
	Key0 KeyCode = KeyCode{code: C.TCODK_0}
	// Key Keypad 0
	KeyKp0 KeyCode = KeyCode{code: C.TCODK_KP0}
	// Key 1
	Key1 KeyCode = KeyCode{code: C.TCODK_1}
	// Key Keypad 1
	KeyKp1 KeyCode = KeyCode{code: C.TCODK_KP1}
	// Key 2
	Key2 KeyCode = KeyCode{code: C.TCODK_2}
	// Key Keypad 2
	KeyKp2 KeyCode = KeyCode{code: C.TCODK_KP2}
	// Key 3
	Key3 KeyCode = KeyCode{code: C.TCODK_3}
	// Key Keypad 3
	KeyKp3 KeyCode = KeyCode{code: C.TCODK_KP3}
	// Key 4
	Key4 KeyCode = KeyCode{code: C.TCODK_4}
	// Key Keypad 4
	KeyKp4 KeyCode = KeyCode{code: C.TCODK_KP4}
	// Key 5
	Key5 KeyCode = KeyCode{code: C.TCODK_5}
	// Key Keypad 5
	KeyKp5 KeyCode = KeyCode{code: C.TCODK_KP5}
	// Key 6
	Key6 KeyCode = KeyCode{code: C.TCODK_6}
	// Key Keypad 6
	KeyKp6 KeyCode = KeyCode{code: C.TCODK_KP6}
	// Key 7
	Key7 KeyCode = KeyCode{code: C.TCODK_7}
	// Key Keypad 7
	KeyKp7 KeyCode = KeyCode{code: C.TCODK_KP7}
	// Key 8
	Key8 KeyCode = KeyCode{code: C.TCODK_8}
	// Key Keypad 8
	KeyKp8 KeyCode = KeyCode{code: C.TCODK_KP8}
	// Key 9
	Key9 KeyCode = KeyCode{code: C.TCODK_9}
	// Key Keypad 9
	KeyKp9 KeyCode = KeyCode{code: C.TCODK_KP9}
)

var (
	// Event Any
	EvAny int = int(C.TCOD_EVENT_ANY)
	// Event None
	EvNone int = int(C.TCOD_EVENT_NONE)
	// Event Finger -
	EvFinger int = int(C.TCOD_EVENT_FINGER)
	// Event Finger Move
	EvFingerMove int = int(C.TCOD_EVENT_FINGER_MOVE)
	// Event Finger Press
	EvFingerPress int = int(C.TCOD_EVENT_FINGER_PRESS)
	// Event Finger Release
	EvFingerRelease int = int(C.TCOD_EVENT_FINGER_RELEASE)
	// Event Key -
	EvKey int = int(C.TCOD_EVENT_KEY)
	// Event Key Press
	EvKeyPress int = int(C.TCOD_EVENT_KEY_PRESS)
	// Event Key Release
	EvKeyRelease int = int(C.TCOD_EVENT_KEY_RELEASE)
	// Event Mouse -
	EvMouse int = int(C.TCOD_EVENT_MOUSE)
	// Event Mouse Move
	EvMouseMove int = int(C.TCOD_EVENT_MOUSE_MOVE)
	// Event Mouse Press
	EvMousePress int = int(C.TCOD_EVENT_MOUSE_PRESS)
	// Event Mouse Release
	EvMouseRelease int = int(C.TCOD_EVENT_MOUSE_RELEASE)
)
