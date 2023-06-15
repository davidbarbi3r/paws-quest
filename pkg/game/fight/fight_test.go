package fight

import (
	"example/paws-quest/pkg/game"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartPlayerTurn_ApplyCurses(t *testing.T) {
	// Créer un GameContext mock
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{
				game.Health:  10,
				game.Stamina: 10,
			},
			Curses: []game.Effect{
				{Field: game.Health, Duration: 2, Amount: 5},
				{Field: game.Stamina, Duration: 3, Amount: 2},
			},
		},
		Destination: &game.Character{
			Parameters: map[game.Field]int{
				game.Health: 10,
			},
		},
	}

	g := &game.Game{
		State:       game.EnemyTurn,
		GameContext: gc,
	}

	// Turn 1
	StartPlayerTurn(gc, g)

	// Turn 1 Verify that Curses are applied
	require.Equal(t, 5, gc.Source.Parameters[game.Health])
	require.Equal(t, 8, gc.Source.Parameters[game.Stamina])

	// Turn 2
	StartPlayerTurn(gc, g)

	// Turn 2 Verify that Curses are applied
	require.Equal(t, 0, gc.Source.Parameters[game.Health])
	require.Equal(t, 6, gc.Source.Parameters[game.Stamina])

	// check if the expired curses are removed
	require.Len(t, gc.Source.Curses, 1)

	// check if the character is dead
	require.True(t, isCharacterDead(gc.Source))

	// check if the game state is set to GameOver
	require.Equal(t, game.GameOver, g.State)
}

func TestHasEnoughStamina(t *testing.T) {
	c := game.Character{
		Parameters: map[game.Field]int{
			game.Stamina: 10,
		},
	}
	require.True(t, hasEnoughStamina(&c, 5))
	require.False(t, hasEnoughStamina(&c, 15))
}

func TestPlayCard_NotEnoughStamina(t *testing.T) {
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{
				game.Stamina: 2,
			},
		},
	}

	g := &game.Game{
		State:       game.PlayerTurn,
		GameContext: gc,
	}

	card := game.Card{
		Cost: 5,
	}

	PlayCard(gc, g, card)

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

	g := &game.Game{
		State:       game.EnemyTurn,
		GameContext: gc,
	}

	gc.Source.Parameters[game.Stamina] = 10
	gc.Destination.Parameters[game.Health] = 10

	card := game.Card{
		Cost:   5,
		Action: game.Attack{Dmg: 5},
	}

	PlayCard(gc, g, card)

	// verify that the character's health as been decreased
	require.Equal(t, 5, gc.Destination.Parameters[game.Health])

	// verify that the character's stamina as been decreased
	require.Equal(t, 5, gc.Source.Parameters[game.Stamina])
}

func TestPlayCard_RemoveCardFromHand(t *testing.T) {
	gc := &game.GameContext{
		Source: &game.Character{
			Parameters: map[game.Field]int{
				game.Health:  10,
				game.Stamina: 10,
			},
			Hand: []game.Card{
				{ID: 1, Action: game.Attack{Dmg: 5}},
				{ID: 2, Action: game.Attack{Dmg: 5}, Cost: 5},
			},
		},
		Destination: &game.Character{
			Parameters: map[game.Field]int{
				game.Health: 10,
			},
		},
	}

	g := &game.Game{
		State:       game.PlayerTurn,
		GameContext: gc,
	}

	PlayCard(gc, g, gc.Source.Hand[1])

	require.Len(t, gc.Source.Hand, 1)
	require.Equal(t, 5, gc.Source.Parameters[game.Stamina])
	require.Equal(t, 1, gc.Source.Hand[0].ID)
}

func TestEndPlayerTurn(t *testing.T) {
	type args struct {
		gc *game.GameContext
		g  *game.Game
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EndPlayerTurn(tt.args.gc, tt.args.g)
		})
	}
}
