package simc_config

import (
	"testing"
)

func Test_CharacterToText(t *testing.T) {
	char := &Character{
		Name: "Capoferro",
		Class: "deathknight",
		Origin: "http://origin",
		Thumbnail: "http://thumb",
		Level: 90,
		Race: "orc",
		Role: "attack",
		Position: "back",
		Professions: "the=profession600/otherprofession600",
		Talents: "some_talents_other_talents",
		Glyphs: "some_glyphs_other_glyphs",
		Spec: "frost"}

	assertEqualString(t, char.ToText(), `deathknight=Capoferro
origin="http://origin"
thumbnail="http://thumb"
level=90
race=orc
role=attack
position=back
professions=the=profession600/otherprofession600
talents=some_talents_other_talents
glyphs=some_glyphs_other_glyphs
spec=frost`)
}
