package tmap

// #include <libtcod.h>
import "C"

type FovFlag int

var (
	Basic       FovFlag = C.FOV_BASIC
	Diamond     FovFlag = C.FOV_DIAMOND
	Shadow      FovFlag = C.FOV_SHADOW
	Restrictive FovFlag = C.FOV_RESTRICTIVE
)
