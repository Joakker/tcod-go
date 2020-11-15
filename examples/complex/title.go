package main

import (
	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

var AnimFlag = false

func TitleScreen(con *tcod.Console) {
	if i := tinput.NewInput(); i.Check() == tinput.EvKeyPress {
		switch i.GetVk() {
		case tinput.KeyText:
			if len(InputBuffer) < WinW-12 {
				InputBuffer += i.GetText()
			}
		case tinput.KeyBackspace:
			if len(InputBuffer) > 0 {
				InputBuffer = InputBuffer[:len(InputBuffer)-1]
			}
		case tinput.KeyEnter:
			InitGame()
			return
		}
	}

	con.PrintFrame(5, 5, WinW-10, 3, false, "What's your name?")
	if len(InputBuffer) < WinW-12 {
		con.SetCellBg(6+len(InputBuffer), 6, tcod.White, tcod.BgLighten)
	}
	con.Print(6, 6, InputBuffer)
	if !AnimFlag {
		AnimFlag = tcod.CreditsEmbed(WinW/2-10, 15, true)
	}
}
