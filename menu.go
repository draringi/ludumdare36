package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

const (
	menuItems    = 5
	oddMenuItems = menuItems % 2
	menuR        = 5
)

const (
	returnID = iota
	loadID
	saveID
	settingsID
	quitID
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
		case returnID:
			g.DeleteView("menu")
			err := g.SetCurrentView("control")
			if err != nil {
				log.Println(err)
			}
			//unpause()
		case quitID:
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

func loadMenu(g *gocui.Gui, v *gocui.View) error {
	log.Println("Loading Menu")
	if v == nil || v.Name() != "menu" {
		log.Println("Creating Menu")
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
