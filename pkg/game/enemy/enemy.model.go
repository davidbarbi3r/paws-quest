package enemy

import (
	"example/paws-quest/pkg/game/card"
	"example/paws-quest/pkg/game/character"
)

type Enemy struct {
	ID int 
	Name string 

	Health int 
	Stamina int 
	Speed int 
	Strength int 
	Agility int 
	Intelligence int 
	Type character.Element 

	Actions []Action 
	ActionsPattern []int 
	Loot Loot 
}

type Loot struct {
	Experience int 
	// Items []Item `json:"items"`
	Cards []card.Card 
}

type Action struct {
	Target string 
	Effect string 
	Amount int 
	Duration int 
}