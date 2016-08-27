package main

type Player struct {
  character *Character
  profession *Profession
  party []*Character
  food uint16
  money uint16
  mana uint16
  attributes Attributes
}

type Attributes struct {
  healingRate float64
  maxMana uint16
}

func (p *Player) Init() {
  // Pre-Profession defaults
  p.food = 300
  p.money = 1000
  p.attributes.healingRate = 1.0
  p.attributes.maxMana = 100
  p.character.Init()
  for _, c := range(p.party) {
    c.Init()
  }
  //Apply Profession modifiers
  p.profession.modPlayer(p)

  //Apply variables based off attributes
  p.mana = p.attributes.maxMana
  p.character.Heal(0)
  for _, c := range(p.party) {
    c.Heal(0)
  }
}
