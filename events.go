package main

import (
	"fmt"
	"math/rand"
)

func triggerBadEvent(s *gamestate) {
	i := rand.Intn(len(badEvents))
	event := badEvents[i]
	str := event(s.Player)
	s.logLock.Lock()
	s.Log = append(s.Log, LogEntry{s.Date, str})
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
	aliveChars := []*Character{p.Character}
	for _, c := range p.Party {
		if c.Alive {
			aliveChars = append(aliveChars, c)
		}
	}
	i := rand.Intn(len(aliveChars))
	hurt := aliveChars[i]
	hurt.Health -= 5
	return fmt.Sprintf("%s tripped and fell, hurting their legs. Took 5 damage.", hurt)
}

func lunPatrol(p *Player) string {
	p.Character.Health -= 25
	for _, c := range p.Party {
		if c.Alive {
			c.Health -= 25
		}
	}
	return "Encountered Lunite Patrol. You successfully fought them off, but took heavy damage. -25 to entire party."
}

func suppliesStolen(p *Player) string {
	foodTaken := 20
	if p.Food == 0 {
		return banditAttack(p)
	} else if p.Food < 20 {
		foodTaken = int(p.Food)
		p.Food = 0
	} else {
		p.Food -= 20
	}
	return fmt.Sprintf("Supplies stolen. %d food lost.", foodTaken)
}

func banditAttack(p *Player) string {
	p.Character.Health -= 15
	for _, c := range p.Party {
		if c.Alive {
			c.Health -= 15
		}
	}
	return "Attacked at night by bandits. 15 damage to entire party."
}

func spellMisfire(p *Player) string {
	p.Mana -= 5
	return "Acidently set off spell. Lose 5 mana."
}
