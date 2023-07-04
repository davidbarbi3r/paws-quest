package card

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"example/paws-quest/pkg/models"
)

func TestShuffleCards(t *testing.T) {
	cardsList := []models.Card{
		{
			ID:          1,
			Name:        "card1",
			Description: "card1",
			Cost:        1,
			Rarity:      models.Common,
		},
		{
			ID:          2,
			Name:        "card2",
			Description: "card2",
			Cost:        2,
			Rarity:      models.Common,
		},
		{
			ID:          3,
			Name:        "card3",
			Description: "card3",
			Cost:        3,
			Rarity:      models.Common,
		},
		{
			ID:          4,
			Name:        "card4",
			Description: "card4",
			Cost:        4,
			Rarity:      models.Common,
		},
		{
			ID:          5,
			Name:        "card5",
			Description: "card5",
			Cost:        5,
			Rarity:      models.Common,
		},
	}

	originalCards := make([]models.Card, len(cardsList))
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

