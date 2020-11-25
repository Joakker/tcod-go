package tmap

import "testing"

func TestBlocked(t *testing.T) {
	// The map starts out blocked
	m := NewMap(5, 5)
	p := m.NewPath(1.0)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			m.Set(i, j, true, true)
		}
	}

	if !p.Compute(1, 1, 2, 2) {
		t.Error("Map is blocked")
	}
}
