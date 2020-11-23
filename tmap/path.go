package tmap

// #include <libtcod.h>
import "C"

type Path struct {
	x1, y1 int
	x2, y2 int
	p      C.TCOD_path_t
}

func (m Map) NewPath(dCost float32) *Path {
	return &Path{p: C.TCOD_path_new_using_map(m.m, C.float(dCost))}
}

func (p *Path) Compute(x1, y1, x2, y2 int) bool {
	p.x1, p.y1, p.x2, p.y2 = x1, y1, x2, y2
	return bool(C.TCOD_path_compute(p.p, C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
}

func (p *Path) GetOrigin() (int, int) {
	return p.x1, p.y1
}

func (p *Path) GetDestination() (int, int) {
	return p.x2, p.y2
}

func (p *Path) Size() int {
	return int(C.TCOD_path_size(p.p))
}

func (p *Path) GetPoint(index int) (int, int) {
	var cx, cy C.int
	C.TCOD_path_get(p.p, C.int(index), &cx, &cy)
	return int(cx), int(cy)
}

func (p *Path) Empty() bool {
	return bool(C.TCOD_path_is_empty(p.p))
}

func (p *Path) Walk(recalculate bool) (bool, int, int) {
	var (
		cx, cy C.int
		x, y   int
	)
	ok := bool(C.TCOD_path_walk(p.p, &cx, &cy, C.bool(recalculate)))

	if !ok {
		x, y = 0, 0
	} else {
		x, y = int(cx), int(cy)
	}
	return ok, x, y
}
