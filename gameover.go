package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func triggerGameover(g *gocui.Gui) error {
	WorldState.gameLoop.kill <- true
	maxX, maxY := g.Size()
	v, err := g.SetView("gameover", maxX/2-17, maxY/2-3, maxX/2+17, maxY/2+3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "You Have Died!"
		fmt.Fprintln(v, "Unfortunately you have died...")
		fmt.Fprintln(v, "Hey, at least it wasn't dysentery")
		fmt.Fprintln(v, "")
		fmt.Fprintln(v, "Press Enter to return\nto the start menu")
		err = g.SetCurrentView("gameover")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func gameoverToStartMenu(g *gocui.Gui, v *gocui.View) error {
	// Wipe state
	WorldState = nil
	// Clean Screen
	g.DeleteView("gameover")
	g.DeleteView("rightStatus")
	g.DeleteView("main")
	g.DeleteView("control")
	// Finally switch to start menu
	g.SetLayout(startMenuLayout)
	return nil
}
