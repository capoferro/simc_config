package simc_config

import (
	"testing"
)

func Test_NewItemFromString(t *testing.T) {
	str := "xifeng_longblade_of_the_titanic_guardian,id=105211,upgrade=2,gems=80str_160mastery_60str,enchant=rune_of_razorice,reforge=hit_exp"

	item := NewItemFromString("off_hand", str)

	assertEqualString(t, item.Label, "xifeng_longblade_of_the_titanic_guardian")
	assertEqualInt(t, item.Id, 105211)
	assertEqualInt(t, item.Upgrade, 2)
	assertEqualStringSlice(t, item.Gems, []string{"80str", "160mastery", "60str"})
	assertEqualString(t, item.Reforge, "hit_exp")
	assertEqualString(t, item.Enchant, "rune_of_razorice")
}

func Test_ItemToText(t *testing.T) {
	str := "xifeng_longblade_of_the_titanic_guardian,id=105211,upgrade=2,gems=80str_160mastery_60str,enchant=rune_of_razorice,reforge=hit_exp"

	item := NewItemFromString("off_hand", str)

	assertEqualString(t, item.ToText(), "off_hand=xifeng_longblade_of_the_titanic_guardian,id=105211,upgrade=2,gems=80str_160mastery_60str,enchant=rune_of_razorice,reforge=hit_exp")
}
