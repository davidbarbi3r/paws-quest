package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActions(t *testing.T) {
	// Prend x pts de vie de la cible et récupère ces points de vie dans la source
	c := &GameContext{
		Source: &Character{
			Parameters: map[Field]int{
				Health: 10,
			},
		},
		Destination: &Character{
			Parameters: map[Field]int{
				Health: 10,
			},
		},
	}
	a := Attack{1, 5}
	a.Do(c)
	require.Equal(t, c.Destination.Parameters[Health], 5)

	// Enlève x pts de vie à la cible et ajoute une Curse à la cible
	c = &GameContext{
		Source: &Character{
			Parameters: map[Field]int{
				Health: 10,
			},
		},
		Destination: &Character{
			Parameters: map[Field]int{
				Health: 10,
			},
		},
	}
	b := &Poison{3, 1}
	b.Do(c)
	require.Equal(t, c.Source.Parameters[Health], 10)
	require.Equal(t, c.Destination.Parameters[Health], 10)
	require.Equal(t, c.Destination.Curses[0].Field, Health)
}