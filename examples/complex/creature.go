package main

import "github.com/Joakker/tcod-go"

// Creature is the base type for any mobile entity, including
// the player themself
type Creature struct {
	X, Y   int
	Char   rune
	Color  tcod.Color
	Name   string
	Hp, Mp Stat
	Scores []int
}

// DrawOnto draws the Creature's sprite to the specified console
func (c Creature) DrawOnto(con tcod.Console) {
	con.SetChar(c.X, c.Y, c.Char)
	con.SetCellFg(c.X, c.Y, c.Color)
}
