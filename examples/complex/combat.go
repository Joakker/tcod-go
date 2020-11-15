package main

var StatNames = []string{
	"STR",
	"DEX",
	"CON",
	"INT",
}

type Stat struct {
	Cur, Max int
}

func NewStat(n int) Stat {
	return Stat{Cur: n, Max: n}
}
