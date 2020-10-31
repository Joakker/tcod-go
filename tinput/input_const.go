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
	KeyKpAdd       KeyCode = KeyCode{code: C.TCODK_KPADD}
	KeyKpSub       KeyCode = KeyCode{code: C.TCODK_KPSUB}
	KeyKpDiv       KeyCode = KeyCode{code: C.TCODK_KPDIV}
	KeyKpMul       KeyCode = KeyCode{code: C.TCODK_KPMUL}
	KeyKpDec       KeyCode = KeyCode{code: C.TCODK_KPDEC}
	KeyKpEnter     KeyCode = KeyCode{code: C.TCODK_KPENTER}
	KeyNumlock     KeyCode = KeyCode{code: C.TCODK_NUMLOCK}
	KeyScrolllock  KeyCode = KeyCode{code: C.TCODK_SCROLLLOCK}
	KeySpace       KeyCode = KeyCode{code: C.TCODK_SPACE}
	KeyText        KeyCode = KeyCode{code: C.TCODK_TEXT}
	KeyChar        KeyCode = KeyCode{code: C.TCODK_CHAR}
	KeyF1          KeyCode = KeyCode{code: C.TCODK_F1}
	KeyF2          KeyCode = KeyCode{code: C.TCODK_F2}
	KeyF3          KeyCode = KeyCode{code: C.TCODK_F3}
	KeyF4          KeyCode = KeyCode{code: C.TCODK_F4}
	KeyF5          KeyCode = KeyCode{code: C.TCODK_F5}
	KeyF6          KeyCode = KeyCode{code: C.TCODK_F6}
	KeyF7          KeyCode = KeyCode{code: C.TCODK_F7}
	KeyF8          KeyCode = KeyCode{code: C.TCODK_F8}
	KeyF9          KeyCode = KeyCode{code: C.TCODK_F9}
	KeyF10         KeyCode = KeyCode{code: C.TCODK_F10}
	KeyF11         KeyCode = KeyCode{code: C.TCODK_F11}
	KeyF12         KeyCode = KeyCode{code: C.TCODK_F12}
	Key0           KeyCode = KeyCode{code: C.TCODK_0}
	KeyKp0         KeyCode = KeyCode{code: C.TCODK_KP0}
	Key1           KeyCode = KeyCode{code: C.TCODK_1}
	KeyKp1         KeyCode = KeyCode{code: C.TCODK_KP1}
	Key2           KeyCode = KeyCode{code: C.TCODK_2}
	KeyKp2         KeyCode = KeyCode{code: C.TCODK_KP2}
	Key3           KeyCode = KeyCode{code: C.TCODK_3}
	KeyKp3         KeyCode = KeyCode{code: C.TCODK_KP3}
	Key4           KeyCode = KeyCode{code: C.TCODK_4}
	KeyKp4         KeyCode = KeyCode{code: C.TCODK_KP4}
	Key5           KeyCode = KeyCode{code: C.TCODK_5}
	KeyKp5         KeyCode = KeyCode{code: C.TCODK_KP5}
	Key6           KeyCode = KeyCode{code: C.TCODK_6}
	KeyKp6         KeyCode = KeyCode{code: C.TCODK_KP6}
	Key7           KeyCode = KeyCode{code: C.TCODK_7}
	KeyKp7         KeyCode = KeyCode{code: C.TCODK_KP7}
	Key8           KeyCode = KeyCode{code: C.TCODK_8}
	KeyKp8         KeyCode = KeyCode{code: C.TCODK_KP8}
	Key9           KeyCode = KeyCode{code: C.TCODK_9}
	KeyKp9         KeyCode = KeyCode{code: C.TCODK_KP9}
)

var (
	EvAny           int32 = int32(C.TCOD_EVENT_ANY)
	EvNone          int32 = int32(C.TCOD_EVENT_NONE)
	EvFinger        int32 = int32(C.TCOD_EVENT_FINGER)
	EvFingerMove    int32 = int32(C.TCOD_EVENT_FINGER_MOVE)
	EvFingerPress   int32 = int32(C.TCOD_EVENT_FINGER_PRESS)
	EvFingerRelease int32 = int32(C.TCOD_EVENT_FINGER_RELEASE)
	EvKey           int32 = int32(C.TCOD_EVENT_KEY)
	EvKeyPress      int32 = int32(C.TCOD_EVENT_KEY_PRESS)
	EvKeyRelease    int32 = int32(C.TCOD_EVENT_KEY_RELEASE)
	EvMouse         int32 = int32(C.TCOD_EVENT_MOUSE)
	EvMouseMove     int32 = int32(C.TCOD_EVENT_MOUSE_MOVE)
	EvMousePress    int32 = int32(C.TCOD_EVENT_MOUSE_PRESS)
	EvMouseRelease  int32 = int32(C.TCOD_EVENT_MOUSE_RELEASE)
)
