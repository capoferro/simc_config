package simc_config

import (
	"strings"
	"regexp"
	"fmt"
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

func (a *Action) ToText() string {
	options := ""
	for k, v := range a.Options {
		options = fmt.Sprintf("%s,%s=%s", options, k, v)
	}
	return fmt.Sprintf("%s%s", a.Label, options)
}
