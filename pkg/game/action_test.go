package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttackAction(t *testing.T) {
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
	require.Equal(t, 5, c.Destination.Parameters[Health])
}

func TestPoisonAction(t *testing.T) {
	// Deal x dmg to the target and apply poison
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
	b := &Poison{3, 1}
	b.Do(c)
	require.Equal(t, 10, c.Source.Parameters[Health])
	require.Equal(t, 10, c.Destination.Parameters[Health])
	require.Equal(t, Health, c.Destination.Curses[0].Field)
}

func TestDodgeAttack(t *testing.T) {
	c := &GameContext{
		Source: &Character{
			Parameters: map[Field]int{
				Health: 10,
			},
		},
		Destination: &Character{
			Parameters: map[Field]int{
				Health: 10,
				Dodge:  100,
			},
		},
	}

	a := Attack{1, 5}
	a.Do(c)

	require.Equal(t, 10, c.Destination.Parameters[Health])
}

func TestCrititalAttack(t *testing.T) {
	c := &GameContext{
		Source: &Character{
			Parameters: map[Field]int{
				Health:         10,
				Critical: 		100,
				CriticalDamage: 50,
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

	require.Equal(t, 3, c.Destination.Parameters[Health])
}