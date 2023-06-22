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
	Rarity Rarity

	Actions []IAction
}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
)

func Shuffle[T int | Card | string](items []T) []T {
	// implement shuffle
	randSeed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSeed)
	r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	return items
}

func CreateDeck (size uint8) []Card {
	// goal 60% common, 30% uncommon, 10% rare
	if size < 20 {
		size = 20
	} else if size > 30 {
		size = 30
	}

	deck := make([]Card, size)

	// todo grab cards from database

	return Shuffle(deck)
}