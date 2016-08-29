package main

import (
	"encoding/json"
	"fmt"
	"github.com/jroimartin/gocui"
	"sync"
)

type LoadSaveMode uint8

const (
	LOADMODE = iota
	SAVEMODE
)

const saveSlots = 10

var loadsaveMode LoadSaveMode

func createLoadSaveView(g *gocui.Gui, mode LoadSaveMode) error {
	loadsaveMode = mode
	maxX, maxY := g.Size()
	v, err := g.SetView("loadsave", maxX/2-14, maxY/2-saveSlots/2, maxX/2+14, maxY/2+saveSlots/2+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		for _, g := range PersistantData.SavedGames {
			if g != nil {
				fmt.Fprintf(v, "%s-%.0fkm\n", g.Player.Character.Name, g.Player.KilometersTravelled)
			} else {
				fmt.Fprintln(v, "Empty")
			}
		}
		err = g.SetCurrentView("loadsave")
		if err != nil {
			return err
		}
	}
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	return nil
}

func loadsaveUp(g *gocui.Gui, v *gocui.View) error {
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

func loadsaveDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < 9 {
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

func DeepCopy(dst, src interface{}) error {
	buffer, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, dst)
}

func loadsaveSelect(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		switch loadsaveMode {
		case LOADMODE:
			if PersistantData.SavedGames[cy] != nil {
				WorldState.gameLoop.kill <- true

				WorldState = new(gamestate)
				err := DeepCopy(WorldState, PersistantData.SavedGames[cy])
				if err != nil {
					return err
				}
				WorldState = PersistantData.SavedGames[cy]
				WorldState.logLock = new(sync.Mutex)
				WorldState.gameLoop = InitiateGameLoop(WorldState, g)
				WorldState.Player.lock = new(sync.Mutex)
				g.DeleteView("loadsave")
				g.SetCurrentView("control")
				// In case we're loading from start screen
				g.SetLayout(layout)
				updateGUI(g)
				WorldState.gameLoop.start()
			}
		case SAVEMODE:
			PersistantData.SavedGames[cy] = new(gamestate)
			DeepCopy(PersistantData.SavedGames[cy], WorldState)
			g.DeleteView("loadsave")
			g.SetCurrentView("control")
			savePersistantToDisk()
			updateGUI(g)
		}
	}
	return nil
}
