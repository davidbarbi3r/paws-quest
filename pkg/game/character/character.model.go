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

type Field int

const (
	Health Field = iota
	Stamina
	Speed
	Strength
	Agility
	Intelligence
	HandSize
)

type Effect struct {
	Field Field
	Duration int
	Amount int
}

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

	Curses []Effect 
	// Buffs []Buff 
	// Items []Item

	IsDead bool
}

