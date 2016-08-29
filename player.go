package main

import "sync"

type Player struct {
	lock                sync.Locker
	Character           *Character
	Profession          *Profession
	Party               []*Character
	Food                uint16
	Money               int
	Mana                uint16
	Speed               Speed
	Rationing           Rationing
	KilometersTravelled float64
	Attributes          Attributes
}

type Attributes struct {
	HealingRate  float64
	MaxMana      uint16
	ManaRecovery uint16
	MovementRate float64
}

func (p *Player) Init() {
	// Pre-Profession defaults
	p.lock = new(sync.Mutex)
	p.Food = 300
	p.Money = 1000
	p.Attributes.HealingRate = .5
	p.Attributes.MovementRate = 1.0
	p.Attributes.MaxMana = 100
	p.Character.Init()
	for _, c := range p.Party {
		c.Init()
	}
	//Apply Profession modifiers
	p.Profession.modPlayer(p)

	//Apply variables based off attributes
	p.Mana = p.Attributes.MaxMana
	p.Character.Heal(0)
	for _, c := range p.Party {
		c.Heal(0)
	}
}

func (p *Player) DailyChecks() {
	aliveCharCount := uint16(1)
	for _, c := range p.Party {
		if c.Alive {
			aliveCharCount++
		}
	}
	notEnoughFood := false
	oldRationing := p.Rationing
	foodNeeded := p.Rationing.AmountEaten(aliveCharCount)
	for foodNeeded > p.Food {
		notEnoughFood = true
		p.Rationing++
		foodNeeded = p.Rationing.AmountEaten(aliveCharCount)
	}
	p.Food -= foodNeeded
	p.Rationing.HungerMod(p.Character)
	p.Character.HungerEffect(p.Attributes.HealingRate)
	p.Character.AliveCheck()
	for _, c := range p.Party {
		if c.Alive {
			p.Rationing.HungerMod(c)
			c.HungerEffect(p.Attributes.HealingRate)
			c.AliveCheck()
		}
	}
	if notEnoughFood {
		WorldState.logLock.Lock()
		WorldState.Log = append(WorldState.Log, NewLogEntry(WorldState.Date, "Not enough food for %v rationing, fell back to %v rationing", oldRationing, p.Rationing))
		WorldState.logLock.Unlock()
		p.Rationing = oldRationing
	}
	//Do anything else that needs to be done
	// Recover Mana, if Magus
	if p.Mana < p.Attributes.MaxMana && p.Attributes.ManaRecovery > 0 {
		p.Mana += p.Attributes.ManaRecovery
		if p.Mana > p.Attributes.MaxMana {
			p.Mana = p.Attributes.MaxMana
		}
	}
	// Final death checks
	p.Character.AliveCheck()
	for _, c := range p.Party {
		c.AliveCheck()
	}
}
