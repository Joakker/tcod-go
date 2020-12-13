package tcod

import "testing"

func TestConversion(t *testing.T) {
	m := map[rune]byte{
		'☻': 2,
		'♥': 3,
	}
	for k, v := range m {
		if c := convertRune(k); c != v {
			t.Errorf("Error converting to CP437: Wanted %d, got %d", v, c)
		}
	}
}
