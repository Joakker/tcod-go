package main

import "github.com/Joakker/tcod-go"

type Enemy struct {
	Creature
}

var EnemyList []*Enemy

func SetupEnemies() {
	EnemyList = append(EnemyList, &Enemy{
		Creature: Creature{
			X: 15, Y: 15, Char: 'e', Color: tcod.Orange,
		},
	})
}

func (e *Enemy) Do() {
	if e.X > player.X {
		e.X--
	} else if e.X < player.X {
		e.X++
	} else if e.Y > player.Y {
		e.Y--
	} else if e.Y < player.Y {
		e.Y++
	}
}
