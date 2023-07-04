package fight

import (
	"example/paws-quest/pkg/action"
	"example/paws-quest/pkg/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartPlayerTurn_ApplyCurses(t *testing.T) {
	// Créer un GameContext mock
	gc := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health:  10,
				models.Stamina: 10,
			},
			Curses: []models.Effect{
				{Field: models.Health, Duration: 2, Amount: 5},
				{Field: models.Stamina, Duration: 3, Amount: 2},
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}

	g := &models.Game{
		State:       models.EnemyTurn,
		GameContext: gc,
	}

	// Turn 1
	StartPlayerTurn(gc, g)

	// Turn 1 Verify that Curses are applied
	require.Equal(t, 5, gc.Source.Parameters[models.Health])
	require.Equal(t, 8, gc.Source.Parameters[models.Stamina])

	// Turn 2
	StartPlayerTurn(gc, g)

	// Turn 2 Verify that Curses are applied
	require.Equal(t, 0, gc.Source.Parameters[models.Health])
	require.Equal(t, 6, gc.Source.Parameters[models.Stamina])

	// check if the expired curses are removed
	require.Len(t, gc.Source.Curses, 1)

	// check if the character is dead
	require.True(t, isCharacterDead(gc.Source))

	// check if the models state is set to GameOver
	require.Equal(t, models.GameOver, g.State)
}

func TestHasEnoughStamina(t *testing.T) {
	c := models.Character{
		Parameters: map[models.Field]int{
			models.Stamina: 10,
		},
	}
	require.True(t, hasEnoughStamina(&c, 5))
	require.False(t, hasEnoughStamina(&c, 15))
}

func TestPlayCard_NotEnoughStamina(t *testing.T) {
	gc := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Stamina: 2,
			},
		},
	}

	g := &models.Game{
		State:       models.PlayerTurn,
		GameContext: gc,
	}

	card := models.Card{
		Cost: 5,
	}

	PlayCard(gc, g, card)

	// verify that the character's stamina is not modified
	require.Equal(t, 2, gc.Source.Parameters[models.Stamina])
}

func TestPlayCard_ApplyCardAction(t *testing.T) {
	gc := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{},
		},
	}

	g := &models.Game{
		State:       models.EnemyTurn,
		GameContext: gc,
	}

	gc.Source.Parameters[models.Stamina] = 10
	gc.Destination.Parameters[models.Health] = 10

	card := models.Card{
		Cost:   5,
		Actions: []models.IAction {
			action.Attack{Duration: 1, Amount: 5},
		},
	}

	PlayCard(gc, g, card)

	// verify that the character's health as been decreased
	require.Equal(t, 5, gc.Destination.Parameters[models.Health])

	// verify that the character's stamina as been decreased
	require.Equal(t, 5, gc.Source.Parameters[models.Stamina])
}

func TestPlayCard_RemoveCardFromHand(t *testing.T) {
	gc := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health:  10,
				models.Stamina: 10,
			},
			Hand: []models.Card{
				{ID: 1, Actions: []models.IAction {action.Attack{Duration: 1, Amount: 5}}},
				{ID: 2, Actions: []models.IAction {action.Attack{Duration: 2, Amount: 5}}, Cost: 5},
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}

	g := &models.Game{
		State:       models.PlayerTurn,
		GameContext: gc,
	}

	PlayCard(gc, g, gc.Source.Hand[1])

	require.Len(t, gc.Source.Hand, 1)
	require.Equal(t, 5, gc.Source.Parameters[models.Stamina])
	require.Equal(t, 1, gc.Source.Hand[0].ID)
}

func TestEndPlayerTurn(t *testing.T) {
	type args struct {
		gc *models.GameContext
		g  *models.Game
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


func TestBuffsLogic (t *testing.T) {
	gc := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health:  10,
				models.Dodge: 10,
			},
			Buffs: []models.Effect{
				{Field: models.Health, Duration: 2, Amount: 5},
				{Field: models.Dodge, Duration: 3, Amount: 2},
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}

	g := &models.Game{
		State:       models.EnemyTurn,
		GameContext: gc,
	}

	// Turn 1
	StartPlayerTurn(gc, g)

	// Turn 1 Verify that Buffs are applied
	require.Equal(t, 15, gc.Source.Parameters[models.Health])
	require.Equal(t, 12, gc.Source.Parameters[models.Dodge])

	EndPlayerTurn(gc, g)

	EnemyTurn(gc, g)

	// Turn 2 Verify that Buffs are applied
	StartPlayerTurn(gc, g)

	require.Equal(t, 20, gc.Source.Parameters[models.Health])
	// Dodge should be applied only once
	require.Equal(t, 2, gc.Source.Buffs[0].Duration)
	require.Equal(t, 12, gc.Source.Parameters[models.Dodge])

	// check if the expired curses are removed
	require.Len(t, gc.Source.Buffs, 1)
}