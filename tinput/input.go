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
