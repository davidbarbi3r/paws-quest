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

func Shuffle[T int | Card | string](items []T) []T {
	// implement shuffle
	randSeed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSeed)
	r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	return items
}