package main

var (
	professionList        []*Profession
	professionListOddness int
	professionListOffset  int
)

type playerModifier func(p *Player)

func doctorMod(p *Player) {
	p.Attributes.HealingRate *= 2
	p.Character.MaxHealth += 50
	for _, c := range p.Party {
		c.MaxHealth += 50
	}
}

func magusMod(p *Player) {
	p.Attributes.MaxMana += 50
	p.Attributes.ManaRecovery = 1
}

func nobleMod(p *Player) {
	p.Money += 5000
}

func quillMod(p *Player) {
	p.Attributes.MovementRate = 1.5
}

type Profession struct {
	Name        string
	description string
	modPlayer   playerModifier
}

func init() {
	professionList = []*Profession{
		&Profession{"Medic", "A Medical Person. Party Members Heal Quicker and everyone starts with more health.", doctorMod},
		&Profession{"Magus", "A Magical Person. Able to cast more powerfull magic, regains mana and has a higher mana capacity.", magusMod},
		&Profession{"Noble", "A Rich Person. Starts with more money. Has no other redeeming features.\nJust a boring, rich, snob...", nobleMod},
		&Profession{"Quillian Rogue", "Pie maker by day, rogue by night. Denies that the nation of Belgium ever existed and claims that Brussels was the capital of France.\nHas a bonus to movement", quillMod},
	}
	professionListOffset = len(professionList) / 2
	professionListOddness = len(professionList) - 2*professionListOffset
}

func (p *Profession) String() string {
	return p.Name
}
