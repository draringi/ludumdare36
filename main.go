package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	activeSettings.Tick = gameSpeedMedium
	init_os()
	g := gocui.NewGui()
	err := g.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = false
	g.SetLayout(initLayout)
	err = setKeyBindings(g)
	if err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	barrier := maxY * 3 / 4
	if v, err := g.SetView("main", 0, 0, maxX*3/4, barrier); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "\x1b[0;31mHello world")
	}
	if v, err := g.SetView("rightStatus", maxX*3/4, 0, maxX-1, barrier); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Status"
		printStatusToView(WorldState, v)
	}
	if v, err := g.SetView("control", 0, barrier, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Log"
		err = g.SetCurrentView("control")
		if err != nil {
			log.Println(err)
		}
		printLogsToView(WorldState.log, v)
	}
	return nil
}

func updateGUI(g *gocui.Gui) error {
	v, err := g.View("control")
	if err != nil {
		return nil
	}
	WorldState.logLock.Lock()
	printLogsToView(WorldState.log, v)
	WorldState.logLock.Unlock()
	v, err = g.View("rightStatus")
	if err != nil {
		return nil
	}
	printStatusToView(WorldState, v)
	return nil
}

func printStatusToView(state *gamestate, v *gocui.View) {
	v.Clear()
	v.SetOrigin(0, 0)
	if state.gameLoop.paused {
		fmt.Fprintln(v, "\x1b[0;31m[PAUSED]\x1b[0;37m")
	}
	fmt.Fprintf(v, "Date: %v\n", state.date)
	fmt.Fprintf(v, "\nParty Leader: %s \n\tHealth: %d/%d\n\tMana: %d/%d\n", state.player.character.name, state.player.character.health, state.player.character.maxHealth, state.player.mana, state.player.attributes.maxMana)
	fmt.Fprintln(v, "\nParty:")
	for _, c := range state.player.party {
		fmt.Fprintf(v, "\t%s ", c.name)
		if c.alive {
			fmt.Fprintf(v, "(%d/%d)\n", c.health, c.maxHealth)
		} else {
			fmt.Fprint(v, "(DEAD)\n")
		}
	}
	fmt.Fprintf(v, "\nMoney: %d coins\nFood: %dlb\n", state.player.money, state.player.food)
}
