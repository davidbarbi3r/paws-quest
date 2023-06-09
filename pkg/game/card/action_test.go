package card

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActions(t *testing.T) {
	// Prend x pts de vie de la cible et récupère ces points de vie dans la source
	c := &gameContext{
		Source: &player{
			Hp: 10,
		},
		Destination: &player{
			Hp: 10,
		},
	}
	a := &leechLife{5}
	a.Do(c)
	require.Equal(t, c.Source.Hp, 15)
	require.Equal(t, c.Destination.Hp, 5)

	// Enlève x pts de vie à la cible et ajoute une Curse à la cible
	c = &gameContext{
		Source: &player{
			Hp: 10,
		},
		Destination: &player{
			Hp: 10,
		},
	}
	b := &poison{5, 1, 3}
	b.Do(c)
	require.Equal(t, c.Source.Hp, 10)
	require.Equal(t, c.Destination.Hp, 5)
	require.Equal(t, c.Destination.Curse[0].Field, Health)
}