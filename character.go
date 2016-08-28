package main

// Generic character type, for both the player, and their party
type Character struct {
	name      string
	health    int
	maxHealth int
	hunger    uint8
	alive     bool
	//status *Status
}

func (c *Character) Init() {
	c.maxHealth = 100
	c.hunger = 10
	c.alive = true
}

func (c *Character) Heal(val int) {
	if val > 0 {
		c.health += val
		if c.health > c.maxHealth {
			c.health = c.maxHealth
		}
	} else {
		c.health = c.maxHealth
	}
}
