package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type Speed uint8

const (
	STOPPED Speed = iota
	SLOW
	STANDARD
	FAST
)

func (s Speed) KilometresPerDay() float64 {
	switch s {
	case STOPPED:
		return 0
	case SLOW:
		return 10
	case STANDARD:
		return 20
	case FAST:
		return 40
	}
	return 0
}

func (s Speed) Risk() float64 {
	switch s {
	case STOPPED:
		return 0.005
	case SLOW:
		return 0.05
	case STANDARD:
		return 0.15
	case FAST:
		return 0.40
	}
	return 1
}

func (s Speed) String() string {
	switch s {
	case STOPPED:
		return "Resting"
	case SLOW:
		return "Slow"
	case STANDARD:
		return "Standard"
	case FAST:
		return "Fast"
	}
	return "Unknown"
}

func createSoundMenu(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("soundMenu", maxX/2-7, maxY/2-3, maxX/2+7, maxY/2+3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		fmt.Fprintln(v, "Current")
		fmt.Fprintln(v, "Rest")
		fmt.Fprintln(v, "Slow")
		fmt.Fprintln(v, "Standard")
		fmt.Fprintln(v, "Fast")
	}
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	return nil
}

var unpauseOnSoundMenuLeave bool

func loadSoundMenu(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != "soundMenu" {
		if !WorldState.gameLoop.paused {
			WorldState.gameLoop.TogglePause()
			unpauseOnSoundMenuLeave = true
		} else {
			unpauseOnSoundMenuLeave = false
		}
		err := createSoundMenu(g)
		if err != nil {
			return err
		}
		err = g.SetCurrentView("soundMenu")
		if err != nil {
			return err
		}
		describeRationing(v, 0)
	}
	return nil
}

func describeRationing(v *gocui.View, s int) {
	v.Clear()
	v.SetOrigin(0, 0)
	switch s {
	case 0: // Current
		fmt.Fprintf(v, "Your current speed setting (%v)", WorldState.player.speed)
	case 1: // Rest
		fmt.Fprint(v, "Rest means you stay where you are and rest. You don't make any progress, and you still consume food\nYou do however avoid any bad events and it allows for your party to heal.")
	case 2: // Slow
		fmt.Fprint(v, "Slow means you travel slower than normal. You don't make as much progress, but you have a lower chance of something bad occuring.")
	case 3: // Standard
		fmt.Fprint(v, "Standard speed, Not too fast, but not too slow either. Faster than Slow, but with less risk than Fast.")
	case 4:
		fmt.Fprint(v, "Fast speed, breakneck even. You go as fast as you possibly can, paying no attention to your surroundings.\nYou make lots of progress, but at a much higher risk. Something will probably go wrong.")
	}
}

func soundUp(g *gocui.Gui, v *gocui.View) error {
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
			describeRationing(help, cy-1)
		}
	}
	return nil
}

func soundDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < 4 {
			if err := v.SetCursor(cx, cy+1); err != nil {
				return err
			}
			help, err := g.View("control")
			if err != nil {
				return nil
			}
			describeRationing(help, cy+1)
		}
	}
	return nil
}

func soundSelect(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		switch cy {
		case 1: // Rest
			WorldState.player.speed = STOPPED
		case 2: // Slow
			WorldState.player.speed = SLOW
		case 3: // Standard
			WorldState.player.speed = STANDARD
		case 4: // Fast
			WorldState.player.speed = FAST
		}
	}
	g.DeleteView("soundMenu")
	err := g.SetCurrentView("control")
	if err != nil {
		return err
	}
	logView, err := g.View("control")
	if err != nil {
		return err
	}
	printLogsToView(WorldState.log, logView)
	statusView, err := g.View("rightStatus")
	if err != nil {
		return err
	}
	printStatusToView(WorldState, statusView)
	if unpauseOnSoundMenuLeave {
		WorldState.gameLoop.TogglePause()
	}
	return nil
}
