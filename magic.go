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
	if p.mana < huntSpellCost {
		return "Not enough mana to hunt"
	}
	p.mana -= huntSpellCost
	foodFound := rand.Intn(huntFoodVariance) + huntFoodBase
	p.food += uint16(foodFound)
	return fmt.Sprintf("Found %d lb of food while hunting.", foodFound)
}

func huntAction(g *gocui.Gui, v *gocui.View) error {
	str := huntForFoodWithMagic(WorldState.player)
	WorldState.logLock.Lock()
	WorldState.log = append(WorldState.log, LogEntry{WorldState.date, str})
	WorldState.logLock.Unlock()
	updateGUI(g)
	return nil
}
