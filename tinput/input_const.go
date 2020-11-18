package tinput

// #cgo pkg-config: libtcod
// #include <libtcod.h>
import "C"

// KeyCode represents the key that the user is triggering
type KeyCode struct {
	code C.TCOD_keycode_t
}

var (

	// KeyEscape is key escape
	KeyEscape KeyCode = KeyCode{code: C.TCODK_ESCAPE}
	// KeyBackspace is key backspace
	KeyBackspace KeyCode = KeyCode{code: C.TCODK_BACKSPACE}
	// KeyTab is key tab
	KeyTab KeyCode = KeyCode{code: C.TCODK_TAB}
	// KeyEnter is key enter
	KeyEnter KeyCode = KeyCode{code: C.TCODK_ENTER}
	// KeyShift is key shift
	KeyShift KeyCode = KeyCode{code: C.TCODK_SHIFT}
	// KeyControl is key control
	KeyControl KeyCode = KeyCode{code: C.TCODK_CONTROL}
	// KeyAlt is key alt
	KeyAlt KeyCode = KeyCode{code: C.TCODK_ALT}
	// KeyPause is key pause
	KeyPause KeyCode = KeyCode{code: C.TCODK_PAUSE}
	// KeyCapslock is key capslock
	KeyCapslock KeyCode = KeyCode{code: C.TCODK_CAPSLOCK}
	// KeyPageup is key pageup
	KeyPageup KeyCode = KeyCode{code: C.TCODK_PAGEUP}
	// KeyPagedown is key pagedown
	KeyPagedown KeyCode = KeyCode{code: C.TCODK_PAGEDOWN}
	// KeyEnd is key end
	KeyEnd KeyCode = KeyCode{code: C.TCODK_END}
	// KeyHome is key home
	KeyHome KeyCode = KeyCode{code: C.TCODK_HOME}
	// KeyUp is key up
	KeyUp KeyCode = KeyCode{code: C.TCODK_UP}
	// KeyLeft is key left
	KeyLeft KeyCode = KeyCode{code: C.TCODK_LEFT}
	// KeyRight is key right
	KeyRight KeyCode = KeyCode{code: C.TCODK_RIGHT}
	// KeyDown is key down
	KeyDown KeyCode = KeyCode{code: C.TCODK_DOWN}
	// KeyPrintscreen is key printscreen
	KeyPrintscreen KeyCode = KeyCode{code: C.TCODK_PRINTSCREEN}
	// KeyInsert is key insert
	KeyInsert KeyCode = KeyCode{code: C.TCODK_INSERT}
	// KeyDelete is key delete
	KeyDelete KeyCode = KeyCode{code: C.TCODK_DELETE}
	// KeyLwin is key lwin
	KeyLwin KeyCode = KeyCode{code: C.TCODK_LWIN}
	// KeyRwin is key rwin
	KeyRwin KeyCode = KeyCode{code: C.TCODK_RWIN}
	// KeyApps is key apps
	KeyApps KeyCode = KeyCode{code: C.TCODK_APPS}
	// KeyKpAdd is key kpAdd
	KeyKpAdd KeyCode = KeyCode{code: C.TCODK_KPADD}
	// KeyKpSub is key kpSub
	KeyKpSub KeyCode = KeyCode{code: C.TCODK_KPSUB}
	// KeyKpDiv is key kpDiv
	KeyKpDiv KeyCode = KeyCode{code: C.TCODK_KPDIV}
	// KeyKpMul is key kpMul
	KeyKpMul KeyCode = KeyCode{code: C.TCODK_KPMUL}
	// KeyKpDec is key kpDec
	KeyKpDec KeyCode = KeyCode{code: C.TCODK_KPDEC}
	// KeyKpEnter is key kpEnter
	KeyKpEnter KeyCode = KeyCode{code: C.TCODK_KPENTER}
	// KeyNumlock is key numlock
	KeyNumlock KeyCode = KeyCode{code: C.TCODK_NUMLOCK}
	// KeyScrolllock is key scrolllock
	KeyScrolllock KeyCode = KeyCode{code: C.TCODK_SCROLLLOCK}
	// KeySpace is key space
	KeySpace KeyCode = KeyCode{code: C.TCODK_SPACE}
	// KeyText is key text
	KeyText KeyCode = KeyCode{code: C.TCODK_TEXT}
	// KeyChar is key char
	KeyChar KeyCode = KeyCode{code: C.TCODK_CHAR}
	// KeyF1 is key function 1
	KeyF1 KeyCode = KeyCode{code: C.TCODK_F1}
	// KeyF2 is key function 2
	KeyF2 KeyCode = KeyCode{code: C.TCODK_F2}
	// KeyF3 is key function 3
	KeyF3 KeyCode = KeyCode{code: C.TCODK_F3}
	// KeyF4 is key function 4
	KeyF4 KeyCode = KeyCode{code: C.TCODK_F4}
	// KeyF5 is key function 5
	KeyF5 KeyCode = KeyCode{code: C.TCODK_F5}
	// KeyF6 is key function 6
	KeyF6 KeyCode = KeyCode{code: C.TCODK_F6}
	// KeyF7 is key function 7
	KeyF7 KeyCode = KeyCode{code: C.TCODK_F7}
	// KeyF8 is key function 8
	KeyF8 KeyCode = KeyCode{code: C.TCODK_F8}
	// KeyF9 is key function 9
	KeyF9 KeyCode = KeyCode{code: C.TCODK_F9}
	// KeyF10 is key function 10
	KeyF10 KeyCode = KeyCode{code: C.TCODK_F10}
	// KeyF11 is key function 11
	KeyF11 KeyCode = KeyCode{code: C.TCODK_F11}
	// KeyF12 is key function 12
	KeyF12 KeyCode = KeyCode{code: C.TCODK_F12}
	// Key0 is key 0
	Key0 KeyCode = KeyCode{code: C.TCODK_0}
	// KeyKp0 is key keypad 0
	KeyKp0 KeyCode = KeyCode{code: C.TCODK_KP0}
	// Key1 is key 1
	Key1 KeyCode = KeyCode{code: C.TCODK_1}
	// KeyKp1 is key keypad 1
	KeyKp1 KeyCode = KeyCode{code: C.TCODK_KP1}
	// Key2 is key 2
	Key2 KeyCode = KeyCode{code: C.TCODK_2}
	// KeyKp2 is key keypad 2
	KeyKp2 KeyCode = KeyCode{code: C.TCODK_KP2}
	// Key3 is key 3
	Key3 KeyCode = KeyCode{code: C.TCODK_3}
	// KeyKp3 is key keypad 3
	KeyKp3 KeyCode = KeyCode{code: C.TCODK_KP3}
	// Key4 is key 4
	Key4 KeyCode = KeyCode{code: C.TCODK_4}
	// KeyKp4 is key keypad 4
	KeyKp4 KeyCode = KeyCode{code: C.TCODK_KP4}
	// Key5 is key 5
	Key5 KeyCode = KeyCode{code: C.TCODK_5}
	// KeyKp5 is key keypad 5
	KeyKp5 KeyCode = KeyCode{code: C.TCODK_KP5}
	// Key6 is key 6
	Key6 KeyCode = KeyCode{code: C.TCODK_6}
	// KeyKp6 is key keypad 6
	KeyKp6 KeyCode = KeyCode{code: C.TCODK_KP6}
	// Key7 is key 7
	Key7 KeyCode = KeyCode{code: C.TCODK_7}
	// KeyKp7 is key keypad 7
	KeyKp7 KeyCode = KeyCode{code: C.TCODK_KP7}
	// Key8 is key 8
	Key8 KeyCode = KeyCode{code: C.TCODK_8}
	// KeyKp8 is key keypad 8
	KeyKp8 KeyCode = KeyCode{code: C.TCODK_KP8}
	// Key9 is key 9
	Key9 KeyCode = KeyCode{code: C.TCODK_9}
	// KeyKp9 is key keypad 9
	KeyKp9 KeyCode = KeyCode{code: C.TCODK_KP9}
)

var (
	// EvAny is any event
	EvAny int = int(C.TCOD_EVENT_ANY)
	// EvNone is no event
	EvNone int = int(C.TCOD_EVENT_NONE)
	// EvFinger Event Finger -
	EvFinger int = int(C.TCOD_EVENT_FINGER)
	// EvFingerMove Event Finger Move
	EvFingerMove int = int(C.TCOD_EVENT_FINGER_MOVE)
	// EvFingerPress Event Finger Press
	EvFingerPress int = int(C.TCOD_EVENT_FINGER_PRESS)
	// EvFingerRelease Event Finger Release
	EvFingerRelease int = int(C.TCOD_EVENT_FINGER_RELEASE)
	// EvKey Event Key -
	EvKey int = int(C.TCOD_EVENT_KEY)
	// EvKeyPress Event Key Press
	EvKeyPress int = int(C.TCOD_EVENT_KEY_PRESS)
	// EvKeyRelease Event Key Release
	EvKeyRelease int = int(C.TCOD_EVENT_KEY_RELEASE)
	// EvMouse Event Mouse -
	EvMouse int = int(C.TCOD_EVENT_MOUSE)
	// EvMouseMove Event Mouse Move
	EvMouseMove int = int(C.TCOD_EVENT_MOUSE_MOVE)
	// EvMousePress Event Mouse Press
	EvMousePress int = int(C.TCOD_EVENT_MOUSE_PRESS)
	// EvMouseRelease Event Mouse Release
	EvMouseRelease int = int(C.TCOD_EVENT_MOUSE_RELEASE)
)
