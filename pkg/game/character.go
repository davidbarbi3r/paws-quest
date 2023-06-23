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
	Health Field = iota // from 10 to 100
	Stamina // from 0 to 10
	Speed // from 0 to 10
	Critical // from 0 to 100
	CriticalDamage // from 0 to 100
	Dodge // from 0 to 100
	HandSize // from 0 to 10
)

type Character struct {
	ID int 
	Name string 
	Type Element
	
	Parameters map[Field]int
	
	Deck []Card
    Hand []Card 
	CardsPatern []int
	Discard []Card

	Curses []Effect 
	Buffs []Effect 
	// Items []Item

	IsDead bool
}
