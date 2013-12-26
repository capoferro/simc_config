package simc_config

import (
	"strconv"
	"strings"
	"regexp"
	"fmt"
)

type Item struct {
	Slot string
	Label string
	Id int
	Upgrade int
	Reforge string
	Gems []string
	Enchant string
}
func NewItemFromString(slot string, str string) *Item {
	item := &Item{ Slot: slot }
	tokens := strings.Split(str, ",")

	re := regexp.MustCompile("=")
	for _, token := range tokens {
		if re.MatchString(token) {
			subtokens := strings.Split(token, "=")
			switch subtokens[0] {
			case "id":
				id, err := strconv.ParseInt(subtokens[1], 10, 0)
				if err != nil {
					panic("Couldn't parse int") // todo message
				}
				item.Id = int(id)
			case "upgrade":
				upgrade, err := strconv.ParseInt(subtokens[1], 10, 0)
				if err != nil {
					panic("Couldn't parse int") // todo message
				}
				item.Upgrade = int(upgrade)
			case "gems":
				item.Gems = gemsFromString(subtokens[1])
			case "reforge":
				item.Reforge = subtokens[1]
			case "enchant":
				item.Enchant = subtokens[1]
			}
		} else {
			item.Label = token
		}
	}

	return item
}

func gemsFromString(str string) []string {
	return strings.Split(str, "_")
}

func (i *Item) gemsToString() string {
	return strings.Join(i.Gems, "_")
}

func (i *Item) ToText() string {
	elements := make([]string, 0)
	if i.Slot != "" && i.Label != "" {
		elements = append(elements, fmt.Sprintf("%s=%s", i.Slot, i.Label))
	}
	if i.Id != 0 {
		elements = append(elements, fmt.Sprintf("id=%d", i.Id))
	}
	if i.Upgrade != 0 {
		elements = append(elements, fmt.Sprintf("upgrade=%d", i.Upgrade))
	}
	if gems := i.gemsToString(); gems != "" {
		elements = append(elements, fmt.Sprintf("gems=%s", gems))
	}
	if i.Enchant != "" {
		elements = append(elements, fmt.Sprintf("enchant=%s", i.Enchant))
	}
	if i.Reforge != "" {
		elements = append(elements, fmt.Sprintf("reforge=%s", i.Reforge))
	}

	return strings.Join(elements, ",")
}
