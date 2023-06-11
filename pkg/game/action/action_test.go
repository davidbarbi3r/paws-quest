package action

import (
	"example/paws-quest/pkg/game/character"
	"example/paws-quest/pkg/game"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActions(t *testing.T) {
	// Prend x pts de vie de la cible et récupère ces points de vie dans la source
	c := &game.GameContext{
		Source: &character.Character{
			Health: 10,
		},
		Destination: &character.Character{
			Health: 10,
		},
	}
	a := &leechLife{5}
	a.Do(c)
	require.Equal(t, c.Source.Health, 15)
	require.Equal(t, c.Destination.Health, 5)

	// Enlève x pts de vie à la cible et ajoute une Curse à la cible
	c = &game.GameContext{
		Source: &character.Character{
			Health: 10,
		},
		Destination: &character.Character{
			Health: 10,
		},
	}
	b := &poison{5, 1, 3}
	b.Do(c)
	require.Equal(t, c.Source.Health, 10)
	require.Equal(t, c.Destination.Health, 5)
	require.Equal(t, c.Destination.Curses[0].Field, character.Health)
}