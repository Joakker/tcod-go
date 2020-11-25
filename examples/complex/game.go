package main

import (
	"os"

	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

var (
	// GameConsole is the part of the screen in which the map is drawn
	GameConsole tcod.Console
	// DataConsole is the part of the screen in which the player data is drawn
	DataConsole tcod.Console
)

// TurnCount keeps track of how many turns have elapsed since the start of
// the game
var TurnCount = uint64(0)

// PrintBgFrame prints the background onto the root console
func PrintBgFrame(con tcod.Console) {
	con.PrintFrame(0, 0, WinW, WinH, false, player.Name)
	x := DataConsole.GetW() + 1
	con.VLine(x, 1, WinH-2, tcod.BgNone)
	con.SetChar(x, 0, tcod.CharTeeS)
	con.SetChar(x, WinH-1, tcod.CharTeeN)
}

// InitGame initializes the main game state
func InitGame() {
	MakePlayer()
	SetupEnemies()
	InputBuffer = ""
	UpdateFunc = GameScreen
	DataConsole = tcod.NewConsole(12, WinH-2)
	GameConsole = tcod.NewConsole(WinW-(DataConsole.GetW()+3), WinH-2)
}

// IsTurn returns true if the key press constitutes a turn
func IsTurn(i tinput.Input) bool {
	switch i.GetVk() {
	case tinput.KeyUp, tinput.KeyDown,
		tinput.KeyLeft, tinput.KeyRight:
		return true
	default:
		return false
	}
}

// GameScreen runs the commands to manage and draw the main game screen
func GameScreen(con tcod.Console) {
	if i := tinput.NewInput(); i.Check() == tinput.EvKeyPress {
		if IsTurn(i) {
			TurnCount++
			for _, e := range EnemyList {
				e.Do()
			}
		}
		switch i.GetVk() {
		case tinput.KeyUp:
			player.Y--
		case tinput.KeyDown:
			player.Y++
		case tinput.KeyLeft:
			player.X--
		case tinput.KeyRight:
			player.X++
		case tinput.KeyEscape:
			os.Exit(0)
		case tinput.KeyBackspace:
			player.Hp.Cur--
			if player.Hp.Cur <= 0 {
				UpdateFunc = DefeatScreen
			}
		case tinput.KeyF1:
			SaveGame()
		}
	}

	DataConsole.Clear()
	GameConsole.Clear()

	PrintBgFrame(con)

	player.DrawDataOnto(DataConsole)
	player.DrawOnto(GameConsole)

	for _, e := range EnemyList {
		e.DrawOnto(GameConsole)
	}

	GameConsole.Blit(0, 0, 0, 0, con, DataConsole.GetW()+2, 1, 1.0, 1.0)
	DataConsole.Blit(0, 0, 0, 0, con, 1, 1, 1.0, 1.0)
}
