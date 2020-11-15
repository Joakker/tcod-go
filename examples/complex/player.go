package main

import (
	"github.com/Joakker/tcod-go"
)

type Player struct {
	Creature
	Money int
}

var player Player

func (p Player) DrawDataOnto(con tcod.Console) {
	con.Print(0, 0, "HP: %3d/%3d", player.Hp.Cur, player.Hp.Max)
	con.Print(0, 1, "MP: %3d/%3d", player.Mp.Cur, player.Mp.Max)

	for i, v := range StatNames {
		con.Print(0, 3+i, "%s - %2d", v, p.Scores[i])
	}
	con.SetChar(1, con.GetH()-2, tcod.CharPound)
	con.Print(2, con.GetH()-2, " %4d", player.Money)
}

func MakePlayer() {
	player = Player{
		Creature: Creature{
			X: 1, Y: 1, Char: '@',
			Color: tcod.Green,
			Name:  InputBuffer,
			Hp:    NewStat(10),
			Mp:    NewStat(0),
		},
	}
	for i := 0; i < len(StatNames); i++ {
		player.Scores = append(player.Scores, 10)
	}
}
