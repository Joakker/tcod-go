# tcod-go

tcod-go is a port of the libtcod library to the Go language

## Installation

Assuming you have the `go` executable in your `PATH` and you have
set up the `GOPATH` environment variable, simply run:

```sh
go get -v "github.com/Joakker/tcod-go"
```

and go should do everything else for you.

## Minimal Program

Copy this into your text editor of choice and `go run` it

```go
package main

import (
    "log"

    "github.com/Joakker/tcod-go"
)

func main() {
    root, err := tcod.InitRoot(80, 50, "The adventures of Go", false, tcod.RenderSDL2)

    if err != nil {
        log.Fatal(err)
    }

    input := tcod.Input{}

    for !tcod.WindowClosed() {
        input.Check()
        root.Clear()
            root.PrintFrame(0, 0, 80, 50, false, "The adventures of Go")
        root.Flush()
    }
}
```

You should get something more or less like this:
![example](images/example.png)
