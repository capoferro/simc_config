package simc_config

import (
	"testing"
	"fmt"
)

func Test_parseLine_Class_1(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("deathknight=\"Capoferro\"")
	if config.Character.Class != "deathknight" {
		t.Error(fmt.Sprintf("Class should be 'deathknight' but was '%s'", config.Character.Class))
	}
	if config.Character.Name != "Capoferro" {
		t.Error(fmt.Sprintf("Name should be 'Capoferro' but was '%s'", config.Character.Name))
	}
}

func Test_parseLine_Class_2(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("deathknight=\"deathknight\"")
	if config.Character.Class != "deathknight" {
		t.Error(fmt.Sprintf("Class should be 'deathknight' but was '%s'", config.Character.Class))
	}
	if config.Character.Name != "deathknight" {
		t.Error(fmt.Sprintf("Name should be 'Capoferro' but was '%s'", config.Character.Name))
	}
}

func Test_parseLine_Empty(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("")
	config.parseLine("deathknight=\"Capoferro\"")
	config.parseLine("")

	if config.Character.Class != "deathknight" {
		t.Error(fmt.Sprintf("Class should be 'deathknight' but was '%s'", config.Character.Class))
	}
	if config.Character.Name != "Capoferro" {
		t.Error(fmt.Sprintf("Name should be 'Capoferro' but was '%s'", config.Character.Name))
	}
}


func Test_parseLine_Garbage(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("garbagelineofgarbage=!%!@#!@#!@%=")	
	// don't raise an error
}

func Test_parseLine_Origin(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("origin=\"http://us.battle.net/wow/en/character/runetotem/Capoferro/advanced\"")

	if config.Character.Origin != "http://us.battle.net/wow/en/character/runetotem/Capoferro/advanced" {
		t.Error(fmt.Sprintf("Origin should be 'http://us.battle.net/wow/en/character/runetotem/Capoferro/advanced' but was '%s'", config.Character.Origin))
	}	
}

func Test_parseLine_Thumbnail(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("thumbnail=\"http://us.battle.net/static-render/us/runetotem/5/94963205-avatar.jpg\"")
	if config.Character.Thumbnail != "http://us.battle.net/static-render/us/runetotem/5/94963205-avatar.jpg" {
		t.Error(fmt.Sprintf("Thumbnail should be 'http://us.battle.net/static-render/us/runetotem/5/94963205-avatar.jpg' but was '%s'", config.Character.Thumbnail))
	}	
}

func Test_parseLine_Level(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("level=90")
	if config.Character.Level != 90 {
		t.Error(fmt.Sprintf("Levelshould be 90 but was '%d'", config.Character.Level))
	}
}

func Test_parseLine_Race(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("race=orc")
	if config.Character.Race != "orc" {
		t.Error(fmt.Sprintf("Race should be 'orc' but was '%s'", config.Character.Race))
	}
}

func Test_parseLine_Role(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("role=attack")
	if config.Character.Role != "attack" {
		t.Error(fmt.Sprintf("Role should be 'attack' but was '%s'", config.Character.Role))
	}
}

func Test_parseLine_Position(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("position=back")
	if config.Character.Position != "back" {
		t.Error(fmt.Sprintf("Position should be 'back' but was '%s'", config.Character.Position))
	}
}

func Test_parseLine_Professions(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("professions=jewelcrafting=600/enchanting=600")
	if config.Character.Professions != "jewelcrafting=600/enchanting=600" {
		t.Error(fmt.Sprintf("Position should be 'jewelcrafting=600/enchanting=600' but was '%s'", config.Character.Professions))
	}
}

func Test_parseLine_Talents(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("talents=http://us.battle.net/wow/en/tool/talent-calculator#dZ!110000")
	if config.Character.Talents != "http://us.battle.net/wow/en/tool/talent-calculator#dZ!110000" {
		t.Error(fmt.Sprintf("Talents should be 'http://us.battle.net/wow/en/tool/talent-calculator#dZ!110000' but was '%s'", config.Character.Talents))
	}
}

func Test_parseLine_Glyphs(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("glyphs=antimagic_shell/loud_horn/regenerative_magic/long_winter/army_of_the_dead/tranquil_grip")
	if config.Character.Glyphs != "antimagic_shell/loud_horn/regenerative_magic/long_winter/army_of_the_dead/tranquil_grip" {
		t.Error(fmt.Sprintf("Glyphs should be 'antimagic_shell/loud_horn/regenerative_magic/long_winter/army_of_the_dead/tranquil_grip' but was '%s'", config.Character.Glyphs))
	}
}

func Test_parseLine_Spec(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("spec=frost")
	if config.Character.Spec != "frost" {
		t.Error(fmt.Sprintf("Spec should be 'frost' but was '%s'", config.Character.Spec))
	}
}

func Test_parseLine_Comment(t *testing.T) {
	config := NewSimcConfig()
	line := "#spec=frost"
	config.parseLine(line)
	if config.Character.Spec == "frost" {
		t.Error(fmt.Sprintf("Spec should not be set by this line: %s", line))
	}
}

func Test_parseLine_Comment_2(t *testing.T) {
	config := NewSimcConfig()
	line := "# Some comment"
	config.parseLine(line)
	// No errors
}

func Test_parseLine_PrecombatAction(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("actions.precombat=flask,type=winters_bite")
	label := config.ActionList.Precombat[0].Label
	if label != "flask" {
		t.Error(fmt.Sprintf("Action should be 'flask' but was '%s'", label))
	}
	action_type := config.ActionList.Precombat[0].Options["type"]
	if action_type != "winters_bite" {
		t.Error(fmt.Sprintf("Action type should be 'winters_bite' but was '%s'", action_type))
	}
}

func Test_parseLine_PrecombatAddAction_2(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("actions.precombat=flask,type=winters_bite")
	config.parseLine("actions.precombat+=/food,type=black_pepper_ribs_and_shrimp")
	label := config.ActionList.Precombat[1].Label
	assertEqualString(t, label, "food")
	action_type := config.ActionList.Precombat[1].Options["type"]
	assertEqualString(t, action_type, "black_pepper_ribs_and_shrimp")
}

func Test_parseLine_PrecombatAddAction_3(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("actions.precombat=flask,type=winters_bite")
	config.parseLine("actions.precombat+=/food,type=black_pepper_ribs_and_shrimp")
	config.parseLine("actions.precombat=flask,type=winters_bite")
	length := len(config.ActionList.Precombat)
	assertEqualInt(t, length, 1)
}

func Test_parseLine_AddActionWithCondition(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("actions.single_target+=/frost_strike,if=buff.killing_machine.react|runic_power>88")
	condition := config.ActionList.SingleTarget[0].Options["if"]
	assertEqualString(t, condition, "buff.killing_machine.react|runic_power>88")
}


func Test_parseLine_Item(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("main_hand=malkoroks_skullcleaver,id=105318,upgrade=2,gems=80str_160mastery_60str,enchant=rune_of_the_fallen_crusader,reforge=hit_mastery")
	item := config.Character.Items[0]
	assertEqualString(t, item.Label, "malkoroks_skullcleaver")
	assertEqualString(t, item.Slot, "main_hand")
}

func Test_dequoteValue(t *testing.T) {
	s := dequoteValue("\"omg\"")
	if s != "omg" {
		t.Error(fmt.Sprintf("Expected 'omg' but got '%s'", s))
	}
}

func assertEqualString(t *testing.T, a string, b string) {
  if a != b {
		t.Error(fmt.Sprintf("'%s' should be '%s'", a, b))
	}
}

func assertEqualInt(t *testing.T, a int, b int) {
	if a != b {
		t.Error(fmt.Sprintf("'%d' should be '%d'", a, b))
	}
}
