package tcod

import "C"

func convertString(s string) *C.char {
	buff := []C.char{}
	for _, c := range s {
		buff = append(buff, C.char(convertRune(c)))
	}
	return &buff[0]
}
