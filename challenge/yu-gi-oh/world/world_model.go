package world

import "strings"

type World struct {
	Id   int
	Name string
	Type string
}

func (w World) BonusIncrement(cardType string) int {
	if strings.ToLower(w.Type) == strings.ToLower(cardType) {
		return 300
	}
	return 0
}
