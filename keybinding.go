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
	err = g.SetKeybinding("startMenu", gocui.KeyArrowDown, gocui.ModNone, startDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("startMenu", gocui.KeyArrowUp, gocui.ModNone, startUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("startMenu", gocui.KeyEnter, gocui.ModNone, startSelect)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("creationInterface", gocui.KeyArrowDown, gocui.ModNone, creationDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("creationInterface", gocui.KeyArrowUp, gocui.ModNone, creationUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("creationInterface", gocui.KeyEnter, gocui.ModNone, creationSelect)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("control", '`', gocui.ModNone, loadMenu)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit)
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
