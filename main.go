package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	err := loadPersistantFromDisk()
	if err != nil {
		log.Println(err)
		PersistantData.GameSettings = new(Settings)
		PersistantData.GameSettings.Tick = gameSpeedMedium
	}
	activeSettings = PersistantData.GameSettings
	init_os()
	g := gocui.NewGui()
	err = g.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SelBgColor = gocui.ColorWhite
	g.SelFgColor = gocui.ColorBlack
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
		fmt.Fprintln(v, "So... I ran out of time...\nImagine pretty ASCII Art here...")
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
		printLogsToView(WorldState.Log, v)
	}
	return nil
}

func updateGUI(g *gocui.Gui) error {
	v, err := g.View("control")
	if err != nil {
		return nil
	}
	WorldState.logLock.Lock()
	printLogsToView(WorldState.Log, v)
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
		fmt.Fprintln(v, "\x1b[0;31m[P]aused\x1b[0;37m")
	}
	fmt.Fprintf(v, "Date: %v\n", state.Date)
	fmt.Fprintf(v, "\nParty Leader: %s \n\tProfession: %v\n\tHealth: %d/%d\n\tMana: %d/%d\n\tHunger Levels: %s\n", state.Player.Character.Name, state.Player.Profession, state.Player.Character.Health, state.Player.Character.MaxHealth, state.Player.Mana, state.Player.Attributes.MaxMana, state.Player.Character.HungerString())
	fmt.Fprintln(v, "\nParty:")
	for _, c := range state.Player.Party {
		fmt.Fprintf(v, "\t%s ", c.Name)
		if c.Alive {
			fmt.Fprintf(v, "(%d/%d)\n", c.Health, c.MaxHealth)
		} else {
			fmt.Fprint(v, "(DEAD)\n")
		}
	}
	fmt.Fprintf(v, "\nMoney: %d coins\nFood: %dlb\n", state.Player.Money, state.Player.Food)
	if currentCity != nil {
		fmt.Fprintf(v, "\nCurrent city: %s", currentCity.name)
	}
	fmt.Fprintf(v, "\nDistance Travelled: %.0fkm\nNext Location: %.0fkm\nRiga->Toledo: %.0fkm\n\n", state.Player.KilometersTravelled,
		distanceToNextLocation(), totalDistanceToTravel())
	fmt.Fprintf(v, "Current [S]peed: %v\n", state.Player.Speed)
	fmt.Fprintf(v, "Current [R]ationing: %v\n", state.Player.Rationing)
	fmt.Fprintln(v, "\nExtra Actions: [H]unt, [~]Menu\n[T]rade (When in a city)")
}
