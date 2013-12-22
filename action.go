package simc_config

import (
	"strings"
	"regexp"
//	"fmt"
)

type Action struct {
	Label string
	Options map[string]string
}

func NewAction() *Action {
	return &Action{ Options: make(map[string]string) }
}

func NewActionFromString(str string) *Action {
	action := NewAction()
	tokens := strings.Split(str, ",")
	
	re := regexp.MustCompile("=")
	for _, token := range tokens {
		if re.MatchString(token) {
			subtokens := strings.Split(token, "=")
			action.Options[subtokens[0]] = subtokens[1]
		} else {
			action.Label = token
		}
	}
	return action
}
