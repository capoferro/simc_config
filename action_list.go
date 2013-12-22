package simc_config

import (
//	"fmt"
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
