package main

var (
	professionList        []*Profession
	professionListOddness int
	professionListOffset  int
)

type playerModifier func(p *Player)

func doctorMod(p *Player) {
	p.attributes.healingRate *= 2
	p.character.maxHealth += 50
	for _, c := range p.party {
		c.maxHealth += 50
	}
}

func magusMod(p *Player) {
	p.attributes.maxMana += 50
}

func nobleMod(p *Player) {
	p.money += 5000
}

func quillMod(p *Player) {
	p.attributes.movementRate = 1.5
}

type Profession struct {
	name        string
	description string
	modPlayer   playerModifier
}

func init() {
	professionList = []*Profession{
		&Profession{"Doctor", "A Medical Person. Party Members Heal Quicker and everyone starts with more health.", doctorMod},
		&Profession{"Magus", "A Magical Person. Able to cast more powerfull magic, regains mana quicker and has a higher mana capacity.", magusMod},
		&Profession{"Noble", "A Rich Person. Starts with more money. Has no other redeeming features.\nJust a boring, rich, snob...", nobleMod},
		&Profession{"Quillian Rogue", "Pie maker by day, rogue by night. Denies that the nation of Belgium ever existed and claims that Brussels was the capital of France.\nHas a bonus to movement", quillMod},
	}
	professionListOffset = len(professionList) / 2
	professionListOddness = len(professionList) - 2*professionListOffset
}

func (p *Profession) String() string {
	return p.name
}
