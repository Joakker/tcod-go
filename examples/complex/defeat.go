package main

import (
	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

func DefeatScreen(con *tcod.Console) {
	if i := tinput.NewInput(); i.Check() == tinput.EvKeyPress &&
		i.GetVk() == tinput.KeyEnter {
		UpdateFunc = TitleScreen
		return
	}
	con.PrintEx(con.GetW()/2, con.GetH()/2,
		tcod.BgNone, tcod.AlignCenter, "Defeated!")
}
