package main

import (
    "log"

    "github.com/Joakker/tcod-go"
    "github.com/Joakker/tcod-go/tinput"
)

func main() {
	/*
		Initialize the main window using the default font file. In the future the
        user should be able to specify other files.

        tcod.InitRoot returns *tcod.Console and error objects
	*/
    root, err := tcod.InitRoot(80, 50, "The adventures of Go", false, tcod.RenderSDL2)

    // If something went wrong, we let the user know and exit
    if err != nil {
        log.Fatal(err)
    }

    // tinput.Input instances are used to survey any kind of input event
    i := tinput.NewInput()

    for !tcod.WindowClosed() {
        // Surveys input events, but doesn't block the program's workflow
        i.Check()
        // Clears the screen
        root.Clear()
            // Prints a frame around the window, with a title in the upper part
            root.PrintFrame(0, 0, 80, 50, false, "The adventures of Go")
        // Renders the changes to the screen
        tcod.Flush()
    }
}
