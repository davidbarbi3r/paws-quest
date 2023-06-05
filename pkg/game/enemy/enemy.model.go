package enemy

import (
	"example/paws-quest/pkg/game/card"
	"example/paws-quest/pkg/game/cat"
)

type Enemy struct {
	ID int `json:"id"`
	Name string `json:"name"`

	Health int `json:"health"`
	// Stamina int `json:"stamina"`
	Speed int `json:"speed"`
	Strength int `json:"strength"`
	Agility int `json:"agility"`
	Intelligence int `json:"intelligence"`
	Type cat.Element `json:"type"`

	Actions []Action `json:"actions"`
	ActionsPattern []int `json:"actionsPattern"`
	Loot Loot `json:"loot"`
}

type Loot struct {
	Experience int `json:"experience"`
	// Items []Item `json:"items"`
	Cards []card.Card `json:"card"`
}

type Action struct {
	Target string `json:"target"`
	Effect string `json:"effect"`
	Amount int `json:"amount"`
	Duration int `json:"duration"`
}