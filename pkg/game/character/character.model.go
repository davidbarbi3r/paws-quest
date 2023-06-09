package character

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

type Character struct {
	ID int 
	Name string 

	Health int
	Stamina int
	Speed int 
	Strength int 
	Agility int 
	Intelligence int 
	Type Element

	Deck []card.Card
    Hand []card.Card 

	// Curses []Curse `json:"curses"`
	// Buffs []Buff `json:"buffs"`
	// Items []Item `json:"objects"`
}

