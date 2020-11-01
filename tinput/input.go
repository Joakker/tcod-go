package tinput

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

/*
	Input contains methods to survey the user's input, and to
	query the type of that input.
*/
type Input struct {
	key   C.TCOD_key_t
	mouse C.TCOD_mouse_t
}

// NewInput returns a new instance of the Input struct
func NewInput() Input {
	return Input{}
}

// Wait pauses the application until the user makes an input event
func (i *Input) Wait(flush bool) int {
	return int(C.TCOD_sys_wait_for_event(
		C.TCOD_EVENT_ANY, &i.key,
		&i.mouse, C.bool(flush),
	))
}

// Check checks whether the user has made an input and moves on
func (i *Input) Check() int {
	return int(C.TCOD_sys_check_for_event(
		C.TCOD_EVENT_ANY, &i.key, &i.mouse,
	))
}

// GetVk returns the KeyCode of the button being triggered
func (i Input) GetVk() KeyCode {
	return KeyCode{code: i.key.vk}
}

/*
	GetC returns the letter corresponding to the key that
	was triggered. It should only be called if calling
	GetVk returns KeyChar
*/
func (i Input) GetC() byte {
	return byte(i.key.c)
}

/*
	KeyPressed returns true if the triggered key is being
	pressed, and false if it's being released.
*/
func (i Input) KeyPressed() bool {
	return bool(i.key.pressed)
}

/*
	GetText returns the current character being typed by
	the user, independent of which keys they are pressing
	to get it. It should only be called if calling GetVk
	returns KeyText
*/
func (i Input) GetText() string {
	return C.GoString(&i.key.text[0])
}

// LAlt returns true if the left alt key is being pressed
func (i Input) LAlt() bool {
	return bool(i.key.lalt)
}

// RAlt returns true if the right alt key is being pressed
func (i Input) RAlt() bool {
	return bool(i.key.ralt)
}

// LCtrl returns true if the left control key is being pressed
func (i Input) LCtrl() bool {
	return bool(i.key.lctrl)
}

// RCtrl returns true if the right control key is being pressed
func (i Input) RCtrl() bool {
	return bool(i.key.rctrl)
}

// Ctrl returns true if either of the control keys is pressed
func (i Input) Ctrl() bool {
	return i.LCtrl() || i.RCtrl()
}

// LMeta returns true if the left meta key is being pressed
func (i Input) LMeta() bool {
	return  bool(i.key.lmeta)
}

// RMeta returns true if the right meta key is being pressed
func (i Input) RMeta() bool {
	return bool(i.key.rmeta)
}

/*
	Meta returns true if either of the meta keys is being pressed.
	The meta key is what's known as the Windows Key on Windows.
*/
func (i Input) Meta() bool {
	return i.LMeta() || i.RMeta()
}

// Shift returns true if either of the shift keys is being pressed
func (i Input) Shift() bool {
	return bool(i.key.shift)
}
