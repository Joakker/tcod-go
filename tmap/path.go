package tmap

// #include <libtcod.h>
import "C"

// Path represents the path a
type Path struct {
	x1, y1 int
	x2, y2 int
	p      C.TCOD_path_t
}

// NewPath returns a new Path from the map data. The cost of traversing
// diagonally is represented by dCost
func (m Map) NewPath(dCost float32) *Path {
	return &Path{p: C.TCOD_path_new_using_map(m.m, C.float(dCost))}
}

// Compute computes the optimal path between (x1, y1) and (x2, y2).
// It returns true if it's possible to walk between those points,
// and false otherwise
func (p *Path) Compute(x1, y1, x2, y2 int) bool {
	p.x1, p.y1, p.x2, p.y2 = x1, y1, x2, y2
	return bool(C.TCOD_path_compute(p.p, C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
}

// GetOrigin returns the x and y coordinates of the path's starting point
func (p *Path) GetOrigin() (x int, y int) {
	return p.x1, p.y1
}

// GetDestination returns the x and y coordinates of the path's
// end point
func (p *Path) GetDestination() (int, int) {
	return p.x2, p.y2
}

// Size returns the length of the path
func (p *Path) Size() int {
	return int(C.TCOD_path_size(p.p))
}

// GetPoint returns the x and y coordinates of the index-th point
// along the path. This does not delete them from the path
func (p *Path) GetPoint(index int) (int, int) {
	var cx, cy C.int
	C.TCOD_path_get(p.p, C.int(index), &cx, &cy)
	return int(cx), int(cy)
}

// Empty returns true if there are no points left
func (p *Path) Empty() bool {
	return bool(C.TCOD_path_is_empty(p.p))
}

// Walk steps along the path and returns the x and y coordinates
// of the next point along it, deleting them from the path. If
// it's not possible to step, ok will be false.
func (p *Path) Walk(recalculate bool) (ok bool, x int, y int) {
	var (
		cx, cy C.int
	)
	ok = bool(C.TCOD_path_walk(p.p, &cx, &cy, C.bool(recalculate)))

	if !ok {
		x, y = 0, 0
	} else {
		x, y = int(cx), int(cy)
	}
	return ok, x, y
}
