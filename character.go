package simc_config

import (
	"strings"
	"fmt"
)

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

func (c *Character) ToText() string {
	s := fmt.Sprintf
	lines := []string{
		s("%s=%s", c.Class, c.Name),
		s("origin=\"%s\"", c.Origin),
		s("thumbnail=\"%s\"", c.Thumbnail),
		s("level=%d", c.Level),
		s("race=%s", c.Race),
		s("role=%s", c.Role),
		s("position=%s", c.Position),
		s("professions=%s", c.Professions),
		s("talents=%s", c.Talents),
		s("glyphs=%s", c.Glyphs),
		s("spec=%s", c.Spec),
		""}

	item_lines := make([]string, len(c.Items))
	for i, item := range c.Items {
		item_lines[i] = item.ToText()
	}

	return strings.Join(append(lines, item_lines...), "\n")
}
