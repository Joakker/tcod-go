// Package main defines a simple roguelike
package main

import (
	"log"

	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

const (
	// WinW is the window's width
	WinW = 60
	// WinH is the window's heigh
	WinH = 40
)

// Player represents the player character
type Player struct {
	X, Y  int
	Char  rune
	Color tcod.Color
}

// CheckInput surveys input from the user and moves the player accordingly
func (p *Player) CheckInput() {
	// tinput.Input instances are used to survey any kind of input event
	if i := tinput.NewInput(); i.Check() == tinput.EvKeyPress {
		switch i.GetVk() {
		case tinput.KeyUp:
			p.Y--
		case tinput.KeyDown:
			p.Y++
		case tinput.KeyLeft:
			p.X--
		case tinput.KeyRight:
			p.X++
		}

		// Keep the player within the bounds of the window frame
		if p.X < 1 {
			p.X = 1
		} else if p.X > WinW-2 {
			p.X = WinW - 2
		}

		if p.Y < 1 {
			p.Y = 1
		} else if p.Y > WinH-2 {
			p.Y = WinH - 2
		}
	}
}

// DrawUpon draws the player onto c as a green '@' sign
func (p Player) DrawUpon(c tcod.Console) {
	c.PutChar(p.X, p.Y, p.Char, tcod.BgNone)
	c.SetCellFg(p.X, p.Y, p.Color)
}

func main() {
	/*
		Initialize the main window using the default font file. In the future the
		user should be able to specify other files.

		tcod.InitRoot returns *tcod.Console and error objects
	*/
	root, err := tcod.InitRoot(WinW, WinH, "The adventures of Go", false, tcod.RenderSDL2)

	// If something went wrong, we let the user know and exit
	if err != nil {
		log.Fatal(err)
	}

	p := Player{X: 5, Y: 5, Char: '@', Color: tcod.Green}

	// Repeat while the window is open
	for !tcod.WindowClosed() {
		// Surveys input events, but doesn't block the program's workflow
		p.CheckInput()
		// Clears the screen
		root.Clear()
		// Prints a frame around the window, with a title in the upper part
		root.PrintFrame(0, 0, WinW, WinH, false, "The adventures of Go")
		// Draws the player as an '@' sign
		p.DrawUpon(root)
		// Renders the changes to the screen
		tcod.Flush()
	}
}
