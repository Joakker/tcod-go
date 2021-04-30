package tmap

//go:generate ../scripts/makeenum -v GPREFIX=Fov -v GTYPE=FovFlag -v OFILE=map_const.go -v CPREFIX=FOV_ ../resources/fov

// #cgo pkg-config: libtcod
// #include <libtcod.h>
import "C"

// Map defines a grid of walkable/transparent tiles
type Map struct {
	w, h int
	m    C.TCOD_map_t
}

// NewMap returns a new instance of the Map struct
// with the given width and height
func NewMap(w, h int) Map {
	return Map{
		w: w, h: h,
		m: C.TCOD_map_new(C.int(w), C.int(h)),
	}
}

// GetW returns the map's width
func (m Map) GetW() int {
	return m.w
}

// GetH returns the map's height
func (m Map) GetH() int {
	return m.h
}

// Set defines wether the cell at (x, y) is transparent
// (FOV functions can pass through it) and walkable
// (Pathfinding functions can pass through it)
func (m Map) Set(x, y int, transp, walk bool) {
	C.TCOD_map_set_properties(m.m, C.int(x), C.int(y),
		C.bool(transp), C.bool(walk))
}

// Clear erases all data from the map
func (m Map) Clear(transp, walk bool) {
	C.TCOD_map_clear(m.m, C.bool(transp), C.bool(walk))
}

// CopyTo copies all data to the destination map
func (m Map) CopyTo(dest Map) {
	C.TCOD_map_copy(m.m, dest.m)
}

// ComputeFOV calculates wether each cell within the radius is visible
func (m Map) ComputeFOV(x, y, radius int, lightWalls bool, algorithm FovFlag) {
	C.TCOD_map_compute_fov(
		m.m, C.int(x), C.int(y), C.int(radius), C.bool(lightWalls),
		C.TCOD_fov_algorithm_t(algorithm),
	)
}

// IsVisible returns true if the cell at position (x, y) is visible
func (m Map) IsVisible(x, y int) bool {
	return bool(C.TCOD_map_is_in_fov(m.m, C.int(x), C.int(y)))
}

// Transparent returns true if the cell at position (x, y) is transparent
func (m Map) Transparent(x, y int) bool {
	return bool(C.TCOD_map_is_transparent(m.m, C.int(x), C.int(y)))
}

// IsWalkable returns true if the cell at position (x, y) is walkable
func (m Map) IsWalkable(x, y int) bool {
    return bool(C.TCOD_map_is_walkable(m.m, C.int(x), C.int(y)))
}

