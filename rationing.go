package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type Rationing uint8

const (
	rationsFILLING Rationing = iota
	rationsSATISFACTORY
	rationsPOOR
	rationsNOFOOD // Not selectable, only when you run out of food
)

const (
	maxHunger    = 10
	minHunger    = -10
	minFedHunger = -5
)

func (r Rationing) String() string {
	switch r {
	case rationsFILLING:
		return "Filling"
	case rationsSATISFACTORY:
		return "Satisfactory"
	case rationsPOOR:
		return "Poor"
	case rationsNOFOOD:
		return "No Food, Starve"
	default:
		return "UNKNOWN"
	}
}

func (r Rationing) Desc() string {
	switch r {
	case rationsFILLING:
		return "Eat well. No one goes hungry, and everyone can have seconds. Supplies might not last long, but at least your party stays healthy."
	case rationsSATISFACTORY:
		return "Eat enough. You might not end a meal feeling filled to the brim, but at least your not hungry anymore. Party maintains itself."
	case rationsPOOR:
		return "Eat barely anything. One meal a day, and a small one at that. Your food will last longer, but health will degenerate."
	default:
		return "UNKNOWN"
	}
}

func (r Rationing) AmountEaten(n uint16) uint16 {
	switch r {
	case rationsFILLING:
		return 3 * n
	case rationsSATISFACTORY:
		return 2 * n
	case rationsPOOR:
		return n
	default:
		return 0
	}
}

func (r Rationing) HungerMod(c *Character) {
	switch r {
	case rationsFILLING:
		if c.Hunger < maxHunger {
			c.Hunger++
		}
	case rationsSATISFACTORY:
		if c.Hunger > 3 {
			c.Hunger--
		} else if c.Hunger < 1 {
			c.Hunger++
		}
	case rationsPOOR:
		if c.Hunger > minFedHunger {
			c.Hunger--
		} else {
			c.Hunger++
		}
	case rationsNOFOOD:
		if c.Hunger > minHunger {
			c.Hunger -= 2
			if c.Hunger < minHunger {
				c.Hunger = minHunger
			}
		}
	}
}

func createrationingMenu(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("rationingMenu", maxX/2-7, maxY/2-3, maxX/2+7, maxY/2+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		fmt.Fprintln(v, "Current")
		fmt.Fprintln(v, rationsFILLING)
		fmt.Fprintln(v, rationsSATISFACTORY)
		fmt.Fprintln(v, rationsPOOR)
	}
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	return nil
}

var unpauseOnrationingMenuLeave bool

func loadrationingMenu(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != "rationingMenu" {
		if !WorldState.gameLoop.paused {
			WorldState.gameLoop.TogglePause()
			unpauseOnrationingMenuLeave = true
		} else {
			unpauseOnrationingMenuLeave = false
		}
		err := createrationingMenu(g)
		if err != nil {
			return err
		}
		err = g.SetCurrentView("rationingMenu")
		if err != nil {
			return err
		}
		describeSpeed(v, 0)
	}
	return nil
}

func describeSpeed(v *gocui.View, s int) {
	v.Clear()
	v.SetOrigin(0, 0)
	switch s {
	case 0: // Current
		fmt.Fprintf(v, "Your current speed setting (%v)", WorldState.Player.Rationing)
	case 1: // Rest
		fmt.Fprint(v, rationsFILLING.Desc())
	case 2: // Slow
		fmt.Fprint(v, rationsSATISFACTORY.Desc())
	case 3: // Standard
		fmt.Fprint(v, rationsPOOR.Desc())
	}
}

func rationingUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy > 0 {
			if err := v.SetCursor(cx, cy-1); err != nil {
				return err
			}
			help, err := g.View("control")
			if err != nil {
				return nil
			}
			describeSpeed(help, cy-1)
		}
	}
	return nil
}

func rationingDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < 3 {
			if err := v.SetCursor(cx, cy+1); err != nil {
				return err
			}
			help, err := g.View("control")
			if err != nil {
				return nil
			}
			describeSpeed(help, cy+1)
		}
	}
	return nil
}

func rationingSelect(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		switch cy {
		case 1: // Rest
			WorldState.Player.Rationing = rationsFILLING
		case 2: // Slow
			WorldState.Player.Rationing = rationsSATISFACTORY
		case 3: // Standard
			WorldState.Player.Rationing = rationsPOOR
		}
	}
	g.DeleteView("rationingMenu")
	err := g.SetCurrentView("control")
	if err != nil {
		return err
	}
	logView, err := g.View("control")
	if err != nil {
		return err
	}
	printLogsToView(WorldState.Log, logView)
	statusView, err := g.View("rightStatus")
	if err != nil {
		return err
	}
	printStatusToView(WorldState, statusView)
	if unpauseOnrationingMenuLeave {
		WorldState.gameLoop.TogglePause()
	}
	return nil
}
