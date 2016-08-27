package main

import (
	"github.com/jroimartin/gocui"
)

func initLayout(g *gocui.Gui) error {
	g.SetLayout(startMenuLayout)
	return nil
}
