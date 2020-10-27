package tinput

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type Input struct {
	key   C.TCOD_key_t
	mouse C.TCOD_mouse_t
}

// Wait pauses the application until the user makes an input event
func (i *Input) Wait(flush bool) int32 {
	return int32(C.TCOD_sys_wait_for_event(
		C.TCOD_EVENT_ANY, &i.key,
		&i.mouse, C.bool(flush),
	))
}

// Check checks whether the user has made an input and moves on
func (i *Input) Check() int32 {
	return int32(C.TCOD_sys_check_for_event(
		C.TCOD_EVENT_ANY, &i.key, &i.mouse,
	))
}

func (i Input) GetVk() KeyCode {
	return KeyCode{code: i.key.vk}
}

func (i Input) GetC() byte {
	return byte(i.key.c)
}

func (i Input) KeyPressed() bool {
	return bool(i.key.pressed)
}

func (i Input) GetText() string {
	return C.GoString(&i.key.text[0])
}

// func KeyPressed(code KeyCode) bool {
// 	return bool(C.TCOD_console_is_key_pressed(code.code))
// }
