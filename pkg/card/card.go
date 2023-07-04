package card

import (
	"math/rand"
	"time"

	"example/paws-quest/pkg/models"
)

func Shuffle[T int | models.Card | string](items []T) []T {
	// implement shuffle
	randSeed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSeed)
	r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	return items
}

func CreateStarterDeck () []models.Card {
	// goal 60% common, 30% uncommon, 10% rare

	deck := []models.Card{
		// 13 common cards
		commonCards[0],
		commonCards[0],
		commonCards[0],
		commonCards[1],
		commonCards[1],
		commonCards[2],
		commonCards[2],
		commonCards[2],
		commonCards[3],
		commonCards[3],
		commonCards[4],
		commonCards[6],
		commonCards[7],
		// 5 uncommon cards
		uncommonCards[2],
		uncommonCards[3],
		uncommonCards[4],
		uncommonCards[0],
		uncommonCards[0],
		// 2 rare cards
		rareCards[0],
		rareCards[1],
	}

	return Shuffle(deck)
}