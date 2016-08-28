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
	err = g.SetKeybinding("soundMenu", gocui.KeyArrowDown, gocui.ModNone, soundDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("soundMenu", gocui.KeyArrowUp, gocui.ModNone, soundUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("soundMenu", gocui.KeyEnter, gocui.ModNone, soundSelect)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("rationingMenu", gocui.KeyArrowDown, gocui.ModNone, rationingDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("rationingMenu", gocui.KeyArrowUp, gocui.ModNone, rationingUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("rationingMenu", gocui.KeyEnter, gocui.ModNone, rationingSelect)
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
	err = g.SetKeybinding("control", 's', gocui.ModNone, loadSoundMenu)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("control", 'r', gocui.ModNone, loadrationingMenu)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("control", 'p', gocui.ModNone, pauseToggleButton)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("control", 'h', gocui.ModNone, huntAction)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("control", 't', gocui.ModNone, openTradeMenu)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("gameover", gocui.KeyEnter, gocui.ModNone, gameoverToStartMenu)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("locationText", gocui.KeyEnter, gocui.ModNone, locationPressEnter)
	if err != nil {
		return err
	}
	// Trade Window time
	err = g.SetKeybinding("tradingView", gocui.KeyArrowDown, gocui.ModNone, tradeDown)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("tradingView", gocui.KeyArrowUp, gocui.ModNone, tradeUp)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("tradingView", gocui.KeyArrowLeft, gocui.ModNone, tradeLeft)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("tradingView", gocui.KeyArrowRight, gocui.ModNone, tradeRight)
	if err != nil {
		return err
	}
	err = g.SetKeybinding("tradingView", gocui.KeyEnter, gocui.ModNone, tradeEnter)
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
