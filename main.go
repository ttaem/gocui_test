package main

import (
	"log"
//	"fmt"

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
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	g.Highlight = true
	width, height := g.Size()
	//version := "1.0.0"
	leftSideWidth := width / 3
	panelSpacing := 1
	optionsTop := height - 2
	statusFilesBoundary := 2
	filesBoundary := 2 * height / 5


	/* CKLEE_TODO, limit view */

	if v, err := g.SetView("main", leftSideWidth+panelSpacing, 0, width-1, optionsTop); err != nil {
					if err != gocui.ErrUnknownView {
									return err
					}
					v.Title = "Main"
					v.FgColor = gocui.ColorGreen
					v.BgColor = gocui.ColorYellow
	}

	if v, err := g.SetView("status", 0, 0, leftSideWidth, statusFilesBoundary); err != nil {
					if err != gocui.ErrUnknownView {
									return err
					}
					v.Title = "Status"
					v.FgColor = gocui.ColorBlue
					v.BgColor = gocui.ColorCyan
	}

	if v, err := g.SetView("files", 0, statusFilesBoundary+panelSpacing, leftSideWidth, filesBoundary); err != nil {
					if err != gocui.ErrUnknownView {
									return err
					}
					v.Title = "Files"
					v.FgColor = gocui.ColorRed
					v.BgColor = gocui.ColorWhite
	}

	if v, err := g.SetView("commits", 0, filesBoundary+panelSpacing, leftSideWidth, optionsTop); err != nil {
					if err != gocui.ErrUnknownView {
									return err
					}
					v.Title = "Commits"
					v.FgColor = gocui.ColorMagenta
					v.BgColor = gocui.ColorWhite
	}


	return nil
}

func quit (g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
