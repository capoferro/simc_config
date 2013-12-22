package simc_config

type Character struct {
	Name string
	Class string
	Origin string
	Thumbnail string
	Level int
	Race string
	Role string
	Position string
	Professions string
	Talents string
	Glyphs string
	Spec string
	Items []*Item
}

func (c *Character) AddItem(slot string, details string) {
	c.Items = append(c.Items, NewItemFromString(slot, details))
}
