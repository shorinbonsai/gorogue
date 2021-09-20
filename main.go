package main

import (
	"log"

	"github.com/Joakker/tcod-go"
	"github.com/Joakker/tcod-go/tinput"
)

func main() {
    root, err := tcod.InitRoot(80, 50, "The adventures of Go", false, tcod.RenderSDL2)

    if err != nil {
        log.Fatal(err)
    }

    i := tinput.NewInput()

    for !tcod.WindowClosed() {
        i.Check()
        root.Clear()
            root.PrintFrame(0, 0, 80, 50, false, "The adventures of Go")
        tcod.Flush()
    }
}