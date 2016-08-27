package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
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
	if v, err := g.SetView("main", 0, 0, maxX-1, barrier); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	if v, err := g.SetView("control", 0, barrier, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
		err = g.SetCurrentView("control")
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
