package main

import (
	"fmt"
	"math/rand"
)

func triggerBadEvent(s *gamestate) {
	i := rand.Intn(len(badEvents))
	event := badEvents[i]
	str := event(s.player)
	s.logLock.Lock()
	s.log = append(s.log, LogEntry{s.date, str})
	s.logLock.Unlock()
}

type BadEvent func(p *Player) string

var badEvents []BadEvent = []BadEvent{
	tripped,
	lunPatrol,
	suppliesStolen,
	spellMisfire,
	banditAttack,
}

func tripped(p *Player) string {
	aliveChars := []*Character{p.character}
	for _, c := range p.party {
		if c.alive {
			aliveChars = append(aliveChars, c)
		}
	}
	i := rand.Intn(len(aliveChars))
	hurt := aliveChars[i]
	hurt.health -= 5
	return fmt.Sprintf("%s tripped and fell, hurting their legs. Took 5 damage.", hurt)
}

func lunPatrol(p *Player) string {
	p.character.health -= 25
	for _, c := range p.party {
		if c.alive {
			c.health -= 25
		}
	}
	return "Encountered Lunite Patrol. You successfully fought them off, but took heavy damage. -25 to entire party."
}

func suppliesStolen(p *Player) string {
	foodTaken := 20
	if p.food == 0 {
		return banditAttack(p)
	} else if p.food < 20 {
		foodTaken = int(p.food)
		p.food = 0
	} else {
		p.food -= 20
	}
	return fmt.Sprintf("Supplies stolen. %d food lost.", foodTaken)
}

func banditAttack(p *Player) string {
	p.character.health -= 15
	for _, c := range p.party {
		if c.alive {
			c.health -= 15
		}
	}
	return "Attacked at night by bandits. 15 damage to entire party."
}

func spellMisfire(p *Player) string {
	p.mana -= 5
	return "Acidently set off spell. Lose 5 mana."
}
