package tmap

// #include <libtcod.h>
import "C"

type FovType int

var (
	FovBasic       FovType = C.FOV_BASIC
	FovDiamond     FovType = C.FOV_DIAMOND
	FovShadow      FovType = C.FOV_SHADOW
	FovRestrictive FovType = C.FOV_RESTRICTIVE
)
