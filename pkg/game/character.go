package game

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

type Character struct {
	ID int 
	Name string 
	Type Element
	
	Parameters map[Field]int
	
	Deck []Card
    Hand []Card 
	Discard []Card

	Curses []Effect 
	// Buffs []Buff 
	// Items []Item

	IsDead bool
}

