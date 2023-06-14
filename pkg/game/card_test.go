package game

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShuffleCards(t *testing.T) {
	cardsList := []Card{
		{
			ID:          1,
			Name:        "card1",
			Description: "card1",
			Cost:        1,
			Rarity:      int(Common),
		},
		{
			ID:          2,
			Name:        "card2",
			Description: "card2",
			Cost:        2,
			Rarity:      int(Common),
		},
		{
			ID:          3,
			Name:        "card3",
			Description: "card3",
			Cost:        3,
			Rarity:      int(Common),
		},
		{
			ID:          4,
			Name:        "card4",
			Description: "card4",
			Cost:        4,
			Rarity:      int(Common),
		},
		{
			ID:          5,
			Name:        "card5",
			Description: "card5",
			Cost:        5,
			Rarity:      int(Common),
		},
	}

	originalCards := make([]Card, len(cardsList))
	copy(originalCards, cardsList)

	Shuffle(cardsList)

	require.Equal(t, len(cardsList), len(originalCards))

	// check if cardsList are shuffled
	require.False(t, reflect.DeepEqual(cardsList, originalCards))
	require.False(t, (
		cardsList[0].ID == originalCards[0].ID) && (
		cardsList[1].ID == originalCards[1].ID) && (
		cardsList[2].ID == originalCards[2].ID) && (
		cardsList[3].ID == originalCards[3].ID) && (
		cardsList[4].ID == originalCards[4].ID))
}

