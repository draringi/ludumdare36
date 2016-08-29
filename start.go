package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"sync"
)

const (
	startmenuItemCount = 4
)

const (
	startStartID = iota
	startLoadID
	startSettingsID
	startQuitID
)

func drawBackground(v *gocui.View) {
	// Insert cool art here
}

func startUp(g *gocui.Gui, v *gocui.View) error {
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

func startDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < startmenuItemCount-1 {
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

func startSelect(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		switch cy {
		case startStartID:
			g.DeleteView("startMenu")
			g.DeleteView("startBackground")
			WorldState = new(gamestate)
			WorldState.logLock = new(sync.Mutex)
			WorldState.Date = startDate
			WorldState.Player = new(Player)
			WorldState.Player.Character = new(Character)
			initCreationScreenSettings()
			g.SetLayout(charCreationLayout)
		case startLoadID:
			return createLoadSaveView(g, LOADMODE)
		case startQuitID:
			return gocui.ErrQuit
		}
	}
	return nil
}

func startMenuLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("startBackground", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		drawBackground(v)
	}
	if v, err := g.SetView("startMenu", maxX/2-7, maxY*3/4-startmenuItemCount/2-1, maxX/2+7, maxY*3/4+startmenuItemCount/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		fmt.Fprintln(v, "Start")
		fmt.Fprintln(v, "Load")
		fmt.Fprintln(v, "Settings")
		fmt.Fprintln(v, "Quit")
		err = g.SetCurrentView("startMenu")
		if err != nil {
			return err
		}
	}
	return nil
}
