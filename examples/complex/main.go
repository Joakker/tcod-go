// Package main defines a roguelike game
package main

import (
	"log"
	"os"

	"github.com/Joakker/tcod-go"
)

const (
	WinW  = 50
	WinH  = 30
	Title = "The Adventures of Go"
)

var (
	UpdateFunc  func(tcod.Console)
	InputBuffer string
)

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
