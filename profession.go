package main

var (
  professionList []*Profession
  professionListOddness int
  professionListOffset int
)

type playerModifier func(p *Player)

func doctorMod(p *Player) {
  p.attributes.healingRate = 2
  p.character.maxHealth += 50
  for _, c := range(p.party) {
    c.maxHealth += 50
  }
}

func magusMod(p *Player) {
  p.attributes.maxMana += 50
}

func nobleMod(p *Player) {
  p.money += 5000
}

type Profession struct {
  name string
  description string
  modPlayer playerModifier
}

func init() {
  professionList = []*Profession{
    &Profession{"Doctor", "A Medical Person. Party Members Heal Quicker and everyone starts with more health.", doctorMod},
    &Profession{"Magus", "A Magical Person. Able to cast more powerfull magic, regains mana quicker and has a higher mana capacity.", magusMod},
    &Profession{"Noble", "A Rich Person. Starts with more money. Has no other redeeming features.\nJust a boring, rich, snob...", nobleMod},
  }
  professionListOffset = len(professionList)/2
  professionListOddness = len(professionList) - 2 * professionListOffset
}
