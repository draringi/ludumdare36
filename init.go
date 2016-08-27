package main

import (
	"github.com/jroimartin/gocui"
)

func initLayout(g *gocui.Gui) error {
     g.SetLayout(layout)
     return nil
}
