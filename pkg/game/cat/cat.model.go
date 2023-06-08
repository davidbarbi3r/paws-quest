package cat

import (
	"example/paws-quest/pkg/game/card"
)

type Element int

const (
	Normal Element = iota
	Fire  
	Water  
	Earth  
	Air 
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

