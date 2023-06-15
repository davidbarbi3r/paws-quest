package fight

import (
	"example/paws-quest/pkg/game"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestHasEnoughStamina(t *testing.T) {
	c := game.Character{
		Parameters: map[game.Field]int{
			game.Stamina: 10,
		},
	}
	if !hasEnoughStamina(&c, 5) {
		require.Fail(t, "Expected true, got false")
	}
	if hasEnoughStamina(&c, 15) {
		require.Fail(t, "Expected false, got true")
	}
}

func TestPlayCard_NotEnoughStamina(t *testing.T) {
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{
				game.Stamina: 2,
			},
		},
	}
	card := game.Card{
		Cost: 5,
	}

	PlayCard(gc, card)

	// verify that the character's stamina is not modified
	require.Equal(t, 2, gc.Source.Parameters[game.Stamina])
}

func TestPlayCard_ApplyCardAction(t *testing.T) {
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{},
		},
		Destination: &game.Character{
			Parameters: map[game.Field]int{},
		},
	}

	gc.Source.Parameters[game.Stamina] = 10
	gc.Destination.Parameters[game.Health] = 10

	card := game.Card{
		Cost:   5,
		Action: game.Attack{Dmg: 5},
	}

	PlayCard(gc, card)

	// verify that the character's health as been decreased
	require.Equal(t, 5, gc.Destination.Parameters[game.Health])

	// verify that the character's stamina as been decreased
	require.Equal(t, 5, gc.Source.Parameters[game.Stamina])
}

func TestPlayCard_RemoveCardFromHand(t *testing.T) {
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{},
			Hand:       []game.Card{
				{ID: 1, Action: game.Attack{Dmg: 5}},
				{ID: 2, Action: game.Attack{Dmg: 5}, Cost: 5},
			},
		},
		Destination: &game.Character{
			Parameters: map[game.Field]int{},
		},
	}
	gc.Source.Parameters[game.Stamina] = 10
	gc.Destination.Parameters[game.Health] = 10

	PlayCard(gc, gc.Source.Hand[1])

	// require.Len(t, gc.Source.Hand, 1)
	require.Equal(t, 5, gc.Source.Parameters[game.Stamina])
	require.Equal(t, 1, gc.Source.Hand[0].ID)
}
