package card

type CardType string

const (
	Attack CardType = "attack"
	Defense CardType = "defense"
	Effect CardType = "effect"
)

type Card struct {
	ID int 
	Name string 
	Description string 
	Type CardType 
	Cost int
	Rarity int
}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
)

