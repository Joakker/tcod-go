// Package main defines a roguelike game
package main

import (
	"log"
	"os"

	"github.com/Joakker/tcod-go"
)

const (
	// WinW is the width of the game window in tiles
	WinW = 50
	// WinH is the height of the game window in tiles
	WinH = 30
	// Title is the title of the game window
	Title = "The Adventures of Go"
)

var (
	// UpdateFunc is the function that will be called every frame
	// to update the game
	UpdateFunc func(tcod.Console)
	// InputBuffer is the buffer in which the text that the
	// player introduces is stored
	InputBuffer string
)

// FileExists is a utility function that returns true
// if there is a file with the specified path
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	root, err := tcod.InitRoot(WinW, WinH, Title, false, tcod.RenderOpenGL)

	if err != nil {
		log.Fatal(err)
	}

	if FileExists("save.json") {
		UpdateFunc = GameScreen
		LoadGame()
	} else {
		UpdateFunc = TitleScreen
	}

	for !tcod.WindowClosed() {
		root.Clear()
		UpdateFunc(root)
		tcod.Flush()
	}
}
