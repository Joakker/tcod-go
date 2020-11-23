/*
Package tcod provides a wrapper around the libtcod C library. Libtcod is a
roguelike making toolkit. It provides utilities for emulating an ascii art
console. It also offers toolkits for common roguelike operations, like bsp
trees, line-of-sight, and pathfinding calculations.

Prerequisites

This library requires you to have the libtcod C library already installed
on your system, as it depends on being able to link the -ltcod library.

Minimal Program

	package main

	import (
		"log"

		"github.com/Joakker/tcod-go"
		"github.com/Joakker/tcod-go/tinput"
	)

	const (
		Width  = 80
		Height = 50
		Title  = "The Adventures of Go"
	)

	func main() {
		// Initialize the game window, also known as the root console
		root, err := tcod.InitRoot(Width, Height, Title, false, tcod.RenderSDL2)

		// Check if we actually have a window, otherwise present the
		// error to the user and exit with code 1
		if err != nil {
			log.Fatal(err)
		}

		// This struct lets us survey window events
		input := tinput.Input{}

		// Continue while the user hasn't requested the window to close
		for !tcod.WindowClosed() {
			// First check for user input
			input.Check()
			// Then wipe clean the game window
			root.Clear()
				// Paint a frame around the window, with a specified title
				// at the top
				root.PrintFrame(0, 0, 80, 50, true, Title)
			// Finally, present the changes to the user
			tcod.Flush()
		}
	}
*/
package tcod
