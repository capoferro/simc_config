package simc_config

import (
	"fmt"
	"strings"
)
type ActionList struct {
	Basic []*Action
	Precombat []*Action
	Aoe []*Action
	SingleTarget []*Action
}

func NewActionList() *ActionList {
	return &ActionList{}
}

func (a *ActionList) AddBasic(action string) {
	a.Basic = append(a.Basic, NewActionFromString(action))
}

func (a *ActionList) AddPrecombat(action string) {
	a.Precombat = append(a.Precombat, NewActionFromString(action))
}

func (a *ActionList) AddAoe(action string) {
	a.Aoe = append(a.Aoe, NewActionFromString(action))
}

func (a *ActionList) AddSingleTarget(action string) {
	a.SingleTarget = append(a.SingleTarget, NewActionFromString(action))
}

func (a *ActionList) InitializeBasic(action string) {
	a.Basic = make([]*Action, 0)
	a.AddBasic(action)
}

func (a *ActionList) InitializePrecombat(action string) {
	a.Precombat = make([]*Action, 0)
	a.AddPrecombat(action)
}

func (a *ActionList) InitializeAoe(action string) {
	a.Aoe = make([]*Action, 0)
	a.AddAoe(action)
}

func (a *ActionList) InitializeSingleTarget(action string) {
	a.SingleTarget = make([]*Action, 0)
	a.AddSingleTarget(action)
}

func (a *ActionList) ToText() string {
	lines := make([]string, 0)
	s := fmt.Sprintf
	c := func(str string) {
		lines = append(lines, str)
	}
	
	if len(a.Precombat) > 0 {
		c(s("actions.precombat=%s", a.Precombat[0].ToText()))
		for _, action := range a.Precombat[1:] {
			c(s("actions.precombat+=/%s", action.ToText()))
		}
	}

	if len(a.Basic) > 0 {
		c(s("actions=%s", a.Basic[0].ToText()))
		for _, action := range a.Basic[1:] {
			c(s("actions+=/%s", action.ToText()))
		}
	}

	if len(a.Aoe) > 0 {
		c(s("actions.aoe=%s", a.Aoe[0].ToText()))
		for _, action := range a.Aoe[1:] {
			c(s("actions.aoe+=/%s", action.ToText()))
		}
	}

	if len(a.SingleTarget) > 0 {
		c(s("actions.single_target=%s", a.SingleTarget[0].ToText()))
		for _, action := range a.SingleTarget[1:] {
			c(s("actions.single_target+=/%s", action.ToText()))
		}
	}

	return strings.Join(lines, "\n")
}
