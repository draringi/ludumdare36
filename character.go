package main

// Generic character type, for both the player, and their party
type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Hunger    int8
	Alive     bool
	//status *Status // Never did implement this...
}

func (c *Character) String() string {
	return c.Name
}

func (c *Character) Init() {
	c.MaxHealth = 100
	c.Hunger = 10
	c.Alive = true
}

func (c *Character) Heal(val int) int {
	if c.Health == c.MaxHealth {
		return 0
	}
	healed := val
	if val > 0 {
		c.Health += val
		if c.Health > c.MaxHealth {
			healed = c.MaxHealth - c.Health
			c.Health = c.MaxHealth
		}
	} else {
		c.Health = c.MaxHealth
	}
	return healed
}

func (c *Character) HungerEffect(healingRate float64) {
	if c.Hunger > 0 {
		// Heal
		healAmmount := int(healingRate * float64(c.Hunger))
		if healAmmount > 0 {
			healed := c.Heal(healAmmount)
			if healed > 0 {
				WorldState.logLock.Lock()
				WorldState.Log = append(WorldState.Log, NewLogEntry(WorldState.Date, "%v recovered %d points of health", c, healed))
				WorldState.logLock.Unlock()
			}
		}
	} else {
		//take damage, unless at hunger == 0
		c.Health += int(c.Hunger)
	}
}

func (c *Character) AliveCheck() {
	if c.Alive && c.Health < 1 {
		c.Alive = false
		WorldState.logLock.Lock()
		WorldState.Log = append(WorldState.Log, NewLogEntry(WorldState.Date, "%v has died", c))
		WorldState.logLock.Unlock()
	}
}

func (c *Character) HungerString() string {
	if c.Hunger == maxHunger {
		return "Very Full"
	} else if c.Hunger > 7 {
		return "Full"
	} else if c.Hunger > 4 {
		return "Well Fed"
	} else if c.Hunger > -1 {
		return "Satisfied"
	} else if c.Hunger > -3 {
		return "Hungry"
	} else if c.Hunger > -6 {
		return "Very Hungry"
	} else if c.Hunger > -8 {
		return "Ravished"
	} else if c.Hunger > -10 {
		return "Starving"
	} else {
		return "At Death's Door"
	}
}
