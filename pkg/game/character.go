package game

type Element int

const (
	Normal Element = iota
	Fire  
	Water  
	Earth  
	Air 
)

type Field string

const (
	Health Field = "health"
	Stamina Field = "stamina"
	Speed Field = "speed"
	Strength Field = "strength"
	Agility Field = "agility"
	Intelligence Field = "intelligence"
	HandSize Field = "handSize"
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
	HandSize int

	Deck []Card
    Hand []Card 
	Discard []Card

	Curses []Effect 
	// Buffs []Buff 
	// Items []Item

	IsDead bool
}

