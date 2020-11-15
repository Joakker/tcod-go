package main

import "github.com/Joakker/tcod-go"

type Creature struct {
	X, Y   int
	Char   byte
	Color  tcod.Color
	Name   string
	Hp, Mp Stat
	Scores []int
}

func (c Creature) DrawOnto(con tcod.Console) {
	con.SetChar(c.X, c.Y, c.Char)
	con.SetCellFg(c.X, c.Y, c.Color)
}
