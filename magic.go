package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"math/rand"
)

const (
	huntSpellCost    = 5
	huntFoodVariance = 50
	huntFoodBase     = 0
)

func huntForFoodWithMagic(p *Player) string {
	if p.Mana < huntSpellCost {
		return "Not enough mana to hunt"
	}
	p.Mana -= huntSpellCost
	foodFound := rand.Intn(huntFoodVariance) + huntFoodBase
	p.Food += uint16(foodFound)
	return fmt.Sprintf("Found %d lb of food while hunting.", foodFound)
}

func huntAction(g *gocui.Gui, v *gocui.View) error {
	str := huntForFoodWithMagic(WorldState.Player)
	WorldState.logLock.Lock()
	WorldState.Log = append(WorldState.Log, LogEntry{WorldState.Date, str})
	WorldState.logLock.Unlock()
	updateGUI(g)
	return nil
}
