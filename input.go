package tcod

// #cgo LDFLAGS: -ltcod
// #include <libtcod.h>
import "C"

type Input struct {
	key   C.TCOD_key_t
	mouse C.TCOD_mouse_t
}


func (i *Input) Wait(flush bool) int32 {
	return int32(C.TCOD_sys_wait_for_event(
		C.TCOD_EVENT_ANY, &i.key,
		&i.mouse, C.bool(flush),
	))
}

func (i *Input) Check() int32 {
	return int32(C.TCOD_sys_check_for_event(
		C.TCOD_EVENT_ANY, &i.key, &i.mouse,
	))
}
