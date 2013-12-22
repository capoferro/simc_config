package simc_config

import (
	"testing"
)

func Test_ToText(t *testing.T) {
	action := &Action{Label: "auto_attack"}
	assertEqualString(t, action.ToText(), "auto_attack")
}

func Test_ToText_WithOptions(t *testing.T) {
	action := &Action{Label: "auto_attack", Options: map[string]string{"if": "some_condition"}}
	assertEqualString(t, action.ToText(), "auto_attack,if=some_condition")
}
