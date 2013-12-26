package simc_config

import (
	"strings"
	"regexp"
	"strconv"
	"fmt"
)

type SimcConfig struct {
	Character *Character
	ActionList *ActionList
}

func NewSimcConfig() *SimcConfig {
	return &SimcConfig{ Character: &Character{}, ActionList: NewActionList() }
}

func NewSimcConfigFromFile(filename string) *SimcConfig {
  return &SimcConfig{}	
}

func (c *SimcConfig) parseLine(line string) {
	tokens := strings.Split(line, "=")
	switch tokens[0] {
	case "warlock",
		"monk",
		"warrior",
		"paladin",
		"mage",
		"druid",
		"rogue",
		"hunter",
		"shaman",
		"priest",
		"deathknight":
		c.Character.Class = tokens[0]
		c.Character.Name = dequoteValue(tokens[1])
	case "origin":
		c.Character.Origin = dequoteValue(tokens[1])
	case "thumbnail":
		c.Character.Thumbnail = dequoteValue(tokens[1])
	case "level":
		level, err := strconv.ParseInt(tokens[1], 10, 0)
		if err == nil {
			c.Character.Level = int(level)
		} else {
			panic(fmt.Sprintf("Level did not parse to an int. Confirm that '%s' is a proper level assignment.", line))
		}
	case "race":
		c.Character.Race = tokens[1]
	case "role":
		c.Character.Role = tokens[1]
	case "position":
		c.Character.Position = tokens[1]
	case "professions":
		c.Character.Professions = strings.Join(tokens[1:], "=")
	case "talents":
		c.Character.Talents = tokens[1]
	case "glyphs":
		c.Character.Glyphs = tokens[1]
	case "spec":
		c.Character.Spec = tokens[1]
	case "actions.precombat", "actions", "actions.aoe", "actions.single_target":
		subtokens := strings.Split(tokens[0], ".")
		action := strings.Join(tokens[1:], "=")
		if len(subtokens) == 2 {
			switch subtokens[1] {
			case "precombat":
				c.ActionList.InitializePrecombat(action)
			case "aoe":
				c.ActionList.InitializeAoe(action)
			case "single_target":
				c.ActionList.InitializeSingleTarget(action)
			}
		} else {
			c.ActionList.InitializeBasic(action)
		}
	case "actions.precombat+", "actions+", "actions.aoe+", "actions.single_target+":
		subtokens := strings.Split(tokens[0], ".")
		action := strings.Join(tokens[1:], "=")
		if len(subtokens) == 2 {
			switch subtokens[1] {
			case "precombat+":
				c.ActionList.AddPrecombat(action[1:])
			case "aoe+":
				c.ActionList.AddAoe(action[1:])
			case "single_target+":
				c.ActionList.AddSingleTarget(action[1:])
			}
		} else {
			c.ActionList.AddBasic(action[1:])
		}
	case "head", "neck", "shoulders", "back", "chest", "shirt", "wrists", "hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand":
		c.Character.AddItem(tokens[0], strings.Join(tokens[1:], "="))
	}
}

func (c *SimcConfig) ToText() string{
	lines := []string{
		"#!/usr/bin/env simc",
		c.Character.ToText(),
		c.ActionList.ToText()}
	
	return strings.Join(lines, "\n")
}

func dequoteValue(quoted string) string {
	re := regexp.MustCompile("\"([^\"]+)\"")
	submatch := re.FindStringSubmatch(quoted)
	return submatch[1]
}

func removeme() {
	fmt.Printf("")
}
