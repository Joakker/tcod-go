package main

import (
	"fmt"

	"github.com/Joakker/tcod-go/tmap"
)

func main() {
	m := tmap.NewMap(50, 50)

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			m.Set(i, j, true, true)
		}
	}

	p := m.NewPath(1.0)
	p.Compute(1, 3, 32, 32)

	for !p.Empty() {
		ok, x, y := p.Walk(true)
		if ok {
			fmt.Println(x, y)
		}
	}
}
