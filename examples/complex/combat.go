package main

// StatNames is the names of the stats in the order they should
// appear both in Player.Stats and on screen
var StatNames = []string{
	"STR",
	"DEX",
	"CON",
	"INT",
}

// Stat represents one of the basic statistics of the player. It has
// current and max values so as to be able to use buffs and debuffs
type Stat struct {
	Cur, Max int
}

// NewStat returns a Stat struct initialized to the appropriate value
func NewStat(n int) Stat {
	return Stat{Cur: n, Max: n}
}
