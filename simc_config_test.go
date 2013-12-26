package simc_config

import (
	"testing"
	"fmt"
	"os"
	"io/ioutil"
	"path"
)

func Test_parseLine_Class_1(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("deathknight=\"Capoferro\"")
	assertEqualString(t, config.Character.Class, "deathknight")
	assertEqualString(t, config.Character.Name, "Capoferro")
}

func Test_parseLine_Class_2(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("deathknight=\"deathknight\"")
	assertEqualString(t, config.Character.Class, "deathknight")
	assertEqualString(t, config.Character.Name, "deathknight")
}

func Test_parseLine_Empty(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("")
	config.parseLine("deathknight=\"Capoferro\"")
	config.parseLine("")

	assertEqualString(t, config.Character.Class, "deathknight")
	assertEqualString(t, config.Character.Name, "Capoferro")
}


func Test_parseLine_Garbage(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("garbagelineofgarbage=!%!@#!@#!@%=")	
	// don't raise an error
}

func Test_parseLine_Origin(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("origin=\"http://us.battle.net/wow/en/character/runetotem/Capoferro/advanced\"")

	assertEqualString(t, config.Character.Origin, "http://us.battle.net/wow/en/character/runetotem/Capoferro/advanced")
}

func Test_parseLine_Thumbnail(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("thumbnail=\"http://us.battle.net/static-render/us/runetotem/5/94963205-avatar.jpg\"")
	assertEqualString(t, config.Character.Thumbnail, "http://us.battle.net/static-render/us/runetotem/5/94963205-avatar.jpg")
}

func Test_parseLine_Level(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("level=90")
	assertEqualInt(t, config.Character.Level, 90)
}

func Test_parseLine_Race(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("race=orc")
	assertEqualString(t, config.Character.Race, "orc")
}

func Test_parseLine_Role(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("role=attack")
	assertEqualString(t, config.Character.Role, "attack")
}

func Test_parseLine_Position(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("position=back")
	assertEqualString(t, config.Character.Position, "back")
}

func Test_parseLine_Professions(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("professions=jewelcrafting=600/enchanting=600")
	assertEqualString(t, config.Character.Professions, "jewelcrafting=600/enchanting=600")
}

func Test_parseLine_Talents(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("talents=http://us.battle.net/wow/en/tool/talent-calculator#dZ!110000")
	assertEqualString(t, config.Character.Talents, "http://us.battle.net/wow/en/tool/talent-calculator#dZ!110000")
}

func Test_parseLine_Glyphs(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("glyphs=antimagic_shell/loud_horn/regenerative_magic/long_winter/army_of_the_dead/tranquil_grip")
	assertEqualString(t, config.Character.Glyphs, "antimagic_shell/loud_horn/regenerative_magic/long_winter/army_of_the_dead/tranquil_grip")
}

func Test_parseLine_Spec(t *testing.T) {
	config := NewSimcConfig()
	config.parseLine("spec=frost")
	assertEqualString(t, config.Character.Spec, "frost")
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
	assertEqualString(t, label, "flask")
	action_type := config.ActionList.Precombat[0].Options["type"]
	assertEqualString(t, action_type, "winters_bite")
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

func Test_simcConfigToText(t *testing.T) {
	config := NewSimcConfig()

	assertEqualString(t, config.ToText(), `#!/usr/bin/env simc

`)
}

func Test_writeFile(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "Test_writeFile")
	defer func() { os.RemoveAll(tempDir) }()

	config := NewSimcConfig()

	config.WriteFile(tempDir)
	content, _ := ioutil.ReadFile(path.Join(tempDir, ".simc"))
	assertEqualString(t, string(content), `#!/usr/bin/env simc

`) // todo handle empty name

}

func Test_dequoteValue(t *testing.T) {
	s := dequoteValue("\"omg\"")
	if s != "omg" {
		t.Error(fmt.Sprintf("Expected 'omg' but got '%s'", s))
	}
}
