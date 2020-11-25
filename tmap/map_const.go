package tmap

// #include <libtcod.h>
import "C"

// FovFlag enum
type FovFlag int

var (
	// FovBasic basic fov algorithm
	FovBasic FovFlag = C.FOV_BASIC
	// FovDiamond Diamond
	FovDiamond FovFlag = C.FOV_DIAMOND
	// FovShadow FOV using recursive shadowcasting
	FovShadow FovFlag = C.FOV_SHADOW
	// FovRestrictive Mingos' Restrictive Precise Angle Shadowcasting (MRPAS)
	FovRestrictive FovFlag = C.FOV_RESTRICTIVE
)
