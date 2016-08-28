package main

import "sync"

type Player struct {
	lock                sync.Locker
	character           *Character
	profession          *Profession
	party               []*Character
	food                uint16
	money               uint16
	mana                uint16
	speed               Speed
	rationing           Rationing
	kilometersTravelled float64
	attributes          Attributes
}

type Attributes struct {
	healingRate  float64
	maxMana      uint16
	movementRate float64
}

func (p *Player) Init() {
	// Pre-Profession defaults
	p.lock = new(sync.Mutex)
	p.food = 300
	p.money = 1000
	p.attributes.healingRate = .5
	p.attributes.movementRate = 1.0
	p.attributes.maxMana = 100
	p.character.Init()
	for _, c := range p.party {
		c.Init()
	}
	//Apply Profession modifiers
	p.profession.modPlayer(p)

	//Apply variables based off attributes
	p.mana = p.attributes.maxMana
	p.character.Heal(0)
	for _, c := range p.party {
		c.Heal(0)
	}
}

func (p *Player) DailyChecks() {
	aliveCharCount := uint16(1)
	for _, c := range p.party {
		if c.alive {
			aliveCharCount++
		}
	}
	notEnoughFood := false
	oldRationing := p.rationing
	foodNeeded := p.rationing.AmountEaten(aliveCharCount)
	for foodNeeded > p.food {
		notEnoughFood = true
		p.rationing++
		foodNeeded = p.rationing.AmountEaten(aliveCharCount)
	}
	p.food -= foodNeeded
	p.rationing.HungerMod(p.character)
	p.character.HungerEffect(p.attributes.healingRate)
	p.character.AliveCheck()
	for _, c := range p.party {
		if c.alive {
			p.rationing.HungerMod(c)
			c.HungerEffect(p.attributes.healingRate)
			c.AliveCheck()
		}
	}
	if notEnoughFood {
		WorldState.logLock.Lock()
		WorldState.log = append(WorldState.log, NewLogEntry(WorldState.date, "Not enough food for %v rationing, fell back to %v rationing", oldRationing, p.rationing))
		WorldState.logLock.Unlock()
		p.rationing = oldRationing
	}
	//Do anything else that needs to be done

	// Final death checks
	p.character.AliveCheck()
	for _, c := range p.party {
		c.AliveCheck()
	}
}
