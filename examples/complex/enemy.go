package main

import "github.com/Joakker/tcod-go"

// Enemy is the base type for any hostile mob
type Enemy struct {
	Creature
}

// EnemyList holds every enemy currently on the level
var EnemyList []*Enemy

// SetupEnemies populates EnemyList with the initial enemies
func SetupEnemies() {
	EnemyList = append(EnemyList, &Enemy{
		Creature: Creature{
			X: 15, Y: 15, Char: 'e', Color: tcod.Orange,
		},
	})
}

// Do makes the enemy perform an action depending on the
// state of the game
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
