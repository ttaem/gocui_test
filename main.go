package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicf(err)
	}

	if err := g.MainLoop(); err != nil {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) err {
	g.Highlight = true
	width, height := g.Size()
	version := "1.0.0"
	leftSideWidth := width / 3

}
