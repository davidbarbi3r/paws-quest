package action

import (
	"testing"

	"example/paws-quest/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestAttackAction(t *testing.T) {
	// Deal x dmg to the target
	c := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}
	a := Attack{1, 5}
	a.Do(c)
	require.Equal(t, 5, c.Destination.Parameters[models.Health])
}

func TestPoisonAction(t *testing.T) {
	// Deal x dmg to the target and apply poison
	c := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}
	b := &Poison{3, 1}
	b.Do(c)
	require.Equal(t, 10, c.Source.Parameters[models.Health])
	require.Equal(t, 10, c.Destination.Parameters[models.Health])
	require.Equal(t, models.Health, c.Destination.Curses[0].Field)
}

func TestDodgeAttack(t *testing.T) {
	c := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
				models.Dodge:  100,
			},
		},
	}

	a := Attack{1, 5}
	a.Do(c)

	require.Equal(t, 10, c.Destination.Parameters[models.Health])
}

func TestCrititalAttack(t *testing.T) {
	c := &models.GameContext{
		Source: &models.Character{
			Parameters: map[models.Field]int{
				models.Health:         10,
				models.Critical:       100,
				models.CriticalDamage: 50,
			},
		},
		Destination: &models.Character{
			Parameters: map[models.Field]int{
				models.Health: 10,
			},
		},
	}

	a := Attack{1, 5}
	a.Do(c)

	require.Equal(t, 3, c.Destination.Parameters[models.Health])
}
