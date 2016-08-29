package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func distanceToNextLocation() float64 {
	i := WorldState.NextLocation
	if i < 0 || i >= len(locationList) {
		return 0
	}
	return locationList[i].distance - WorldState.Player.KilometersTravelled
}

func totalDistanceToTravel() float64 {
	return locationList[len(locationList)-1].distance
}

func finalDestionation() string {
	return locationList[len(locationList)-1].name
}

type Location struct {
	distance    float64
	name        string
	description string
}

func (l *Location) String() string {
	return l.name
}

var locationList []*Location = []*Location{
	&Location{0, "Riga", "Riga, the last free human city on the Baltic.\nWhile you once called this place home, you realized the city won't last much longer against\nthe threat of the vamperic Lun Princes of the Baltic coast.\nYou've heard rumours of a safe city over 3000km away, on the Mediterranean,\nrich in the ancient technologies from before the breaking.\nYou gather your belongings, and set out for your promised land with your companions."},
	&Location{661, "Warsaw", "Warsaw, your first stop on the trip to Toledo"},
	&Location{1223, "Berlin", "Berlin. Once a great city, Berlin is now a shadow of its former self.\nThe city took a lot of damage in the wars leading up to the breaking, and now it is just a collection of ruined buildings.\nYou find some merchants selling wares near the Brandenburg Gate."},
	&Location{1473, "Hanover", "Hanover. Once the seat of an electorship of the HRE,\nit is now a small duchy, trapped between the Lun princes to the north,\nand the expanding Restored Kingdom of Bavaria to the south."},
	&Location{1761, "Cologne", "Placeholder Text..."},
	&Location{1983, "Brussels", "Placeholder Text..."},
	&Location{2242, "Paris", "Placeholder Text..."},
	&Location{2821, "Bordeaux", "Placeholder Text..."},
	&Location{3059, "San Sebastian", "Placeholder Text..."},
	&Location{3512, "Madrid", "Madrid, the gateway to Toledo.\nWhile the city is mainly deserted, there is still a military outpost, and merchants nearby."},
	&Location{3578, "Toledo", "Toledo, and its underground city. This is your promised land! Congratulations on make it."},
}

var (
	riga          = locationList[0]
	warsaw        = locationList[1]
	berlin        = locationList[2]
	hanover       = locationList[3]
	cologne       = locationList[4]
	brussels      = locationList[5]
	paris         = locationList[6]
	bordeaux      = locationList[7]
	san_sebastian = locationList[8]
	madrid        = locationList[9]
)

var currentCity *Location

func locationCheck(state *gamestate, g *gocui.Gui, eventThread bool) {
	target := locationList[state.NextLocation]
	if state.Player.KilometersTravelled >= target.distance {
		state.gameLoop.SetPause(true)
		// Make sure the player "stopped" in the city
		state.Player.KilometersTravelled = target.distance
		currentCity = target
		state.logLock.Lock()
		state.Log = append(state.Log, NewLogEntry(state.Date, "Arrived in %s", target.name))
		state.logLock.Unlock()
		state.NextLocation++
		if eventThread {
			g.SetLayout(locationLayout)
		} else {
			g.Execute(setLocationLayout)
		}
	}
}

func setLocationLayout(g *gocui.Gui) error {
	g.SetLayout(locationLayout)
	return nil
}

func locationLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("locationBackground", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		drawBackground(v)
	}
	if v, err := g.SetView("locationText", maxX/6, maxY*2/3, maxX*5/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = currentCity.name
		fmt.Fprintln(v, currentCity.description)
		fmt.Fprint(v, "\nPress Enter to Continue")
		err = g.SetCurrentView("locationText")
		if err != nil {
			return err
		}
	}
	return nil
}

func locationPressEnter(g *gocui.Gui, v *gocui.View) error {
	// Clean Screen
	g.DeleteView("locationBackground")
	g.DeleteView("locationBackground")
	// Finally switch to start menu
	g.SetLayout(layout)
	return nil
}
