package tmap

// #include <libtcod.h>
import "C"

type FovFlag int

var (
	FovBasic       FovFlag = C.FOV_BASIC
	FovDiamond     FovFlag = C.FOV_DIAMOND
	FovShadow      FovFlag = C.FOV_SHADOW
	FovRestrictive FovFlag = C.FOV_RESTRICTIVE
)
