package game

import (
	"math/rand"
	"time"
)

type Card struct {
	ID int 
	Name string 
	Description string 
	Cost int
	Rarity int

	Action IAction
}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
)

func ShuffleCards(cards []Card) []Card {
	// implement shuffle
	randSeed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSeed)
	r.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}