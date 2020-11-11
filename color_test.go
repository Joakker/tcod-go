package tcod

import "testing"

func TestAdd(t *testing.T) {
	a := Color{ R: 255, G: 0, B: 0 }
	b := Color{ R: 1, G: 0, B: 0 }
	c := a.Add(b)
	if a != c {
		t.Errorf("Color.Add() failed, expected %v, got %v", a, c)
	}
}

func TestSub(t *testing.T) {
	a := Color{ R: 0, G: 0, B: 0 }
	b := Color{ R: 1, G: 0, B: 0 }
	c := a.Sub(b)
	if a != c {
		t.Errorf("Color.Sub() failed, expected %v, got %v", a, c)
	}
}
