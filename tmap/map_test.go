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

func TestIsWalkable(t *testing.T) {
    m := NewMap(5, 5)
    
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            // Just to populate map with random flags
            walkable := 2%(i+j+1) != 0
            m.Set(i, j, true, walkable)
        }
    }

    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            expectedWalkable := 2%(i+j+1) != 0
            actualWalkable := m.IsWalkable(i, j)

            if(actualWalkable != expectedWalkable) {
                t.Error("Bad Walkable Flag")
            }
        }
    }
}
