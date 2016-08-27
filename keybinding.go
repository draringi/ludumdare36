package main

import (
	"github.com/jroimartin/gocui"
)

func setKeyBindings(g *gocui.Gui) error {
	err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, menuDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, menuUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, menuSelect)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("", '`', gocui.ModNone, loadMenu)
	if err != nil {
		return err
	}
	return nil
}
