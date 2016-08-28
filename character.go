package main

// Generic character type, for both the player, and their party
type Character struct {
	name      string
	health    int
	maxHealth int
	hunger    int8
	alive     bool
	//status *Status
}

func (c *Character) String() string {
	return c.name
}

func (c *Character) Init() {
	c.maxHealth = 100
	c.hunger = 10
	c.alive = true
}

func (c *Character) Heal(val int) int {
	if c.health == c.maxHealth {
		return 0
	}
	healed := val
	if val > 0 {
		c.health += val
		if c.health > c.maxHealth {
			healed = c.maxHealth - c.health
			c.health = c.maxHealth
		}
	} else {
		c.health = c.maxHealth
	}
	return healed
}

func (c *Character) HungerEffect(healingRate float64) {
	if c.hunger > 0 {
		// Heal
		healAmmount := int(healingRate * float64(c.hunger))
		if healAmmount > 0 {
			healed := c.Heal(healAmmount)
			if healed > 0 {
				WorldState.logLock.Lock()
				WorldState.log = append(WorldState.log, NewLogEntry(WorldState.date, "%v recovered %d points of health", c, healed))
				WorldState.logLock.Unlock()
			}
		}
	} else {
		//take damage, unless at hunger == 0
		c.health += int(c.hunger)
	}
}

func (c *Character) AliveCheck() {
	if c.alive && c.health < 1 {
		c.alive = false
		WorldState.logLock.Lock()
		WorldState.log = append(WorldState.log, NewLogEntry(WorldState.date, "%v has died", c))
		WorldState.logLock.Unlock()
	}
}

func (c *Character) HungerString() string {
	if c.hunger == maxHunger {
		return "Very Full"
	} else if c.hunger > 7 {
		return "Full"
	} else if c.hunger > 4 {
		return "Well Fed"
	} else if c.hunger > -1 {
		return "Satisfied"
	} else if c.hunger > -3 {
		return "Hungry"
	} else if c.hunger > -6 {
		return "Very Hungry"
	} else if c.hunger > -8 {
		return "Ravished"
	} else if c.hunger > -10 {
		return "Starving"
	} else {
		return "At Death's Door"
	}
}
