package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	menuItems    = 5
	oddMenuItems = menuItems % 2
	menuR        = 5
)

const (
	menuReturnID = iota
	menuLoadID
	menuSaveID
	menuSettingsID
	menuQuitID
)

func menuUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if cy > 0 {
			if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
				if err := v.SetOrigin(ox, oy-1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func menuDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < menuItems-1 {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func menuSelect(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		switch cy {
		case menuReturnID:
			g.DeleteView("menu")
			err := g.SetCurrentView("control")
			if err != nil {
				return err
			}
			if unpauseOnMenuLeave {
				WorldState.gameLoop.TogglePause()
			}
		case menuQuitID:
			return gocui.ErrQuit
		}
	}
	return nil
}

func createMenu(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("menu", maxX/2-menuR, maxY/2-menuItems/2, maxX/2+menuR, maxY/2+menuItems/2+oddMenuItems+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		fmt.Fprintln(v, "Return")
		fmt.Fprintln(v, "Load")
		fmt.Fprintln(v, "Save")
		fmt.Fprintln(v, "Settings")
		fmt.Fprintln(v, "Quit")
	}
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	return nil
}

var unpauseOnMenuLeave bool

func loadMenu(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != "menu" {
		if !WorldState.gameLoop.paused {
			WorldState.gameLoop.TogglePause()
			unpauseOnMenuLeave = true
		} else {
			unpauseOnMenuLeave = false
		}
		err := createMenu(g)
		if err != nil {
			return err
		}
		err = g.SetCurrentView("menu")
		if err != nil {
			return err
		}
	}
	return nil
}
