package simc_config

import (
	"testing"
)

func Test_ActionListToText(t *testing.T) {
	list := &ActionList{
		Precombat: []*Action{
			&Action{Label: "army_of_the_dead"},
			&Action{
				Label: "potion", 
				Options: map[string]string{"type": "potion_of_mogu_power"}}},
		Basic: []*Action{
			&Action{
				Label: "auto_attack",
				Options: map[string]string{"synchronize_swing": "1"}}}}

	assertEqualString(t, list.ToText(), `actions.precombat=army_of_the_dead
actions.precombat+=/potion,type=potion_of_mogu_power
actions=auto_attack,synchronize_swing=1`)
}
