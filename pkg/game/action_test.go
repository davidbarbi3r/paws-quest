package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActions(t *testing.T) {
	// Deal x dmg to the target
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

	// Deal x dmg to the target and apply poison
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