package cat

import (
	"example/paws-quest/pkg/game/card"
)

type Element string

const (
	Normal Element = "normal"
	Fire Element = "fire"
	Water Element = "water"
	Earth Element = "earth"
	Air Element = "air"
)

type Cat struct {
	ID int `json:"id"`
	Name string `json:"name"`

	Health int `json:"health"`
	Stamina int `json:"stamina"`
	Speed int `json:"speed"`
	Strength int `json:"strength"`
	Agility int `json:"agility"`
	Intelligence int `json:"intelligence"`
	Type Element `json:"type"`	

	// Curses []Curse `json:"curses"`
	// Buffs []Buff `json:"buffs"`
	Deck []card.Card `json:"deck"`
	Hand []card.Card `json:"hand"`
	// Items []Item `json:"objects"`
}
