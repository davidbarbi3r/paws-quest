package action

import (
	"math/rand"
	"time"

	"example/paws-quest/pkg/models"
)

type Poison struct {
	Duration int
	Amount   int
}

type Heal struct {
	Duration int
	Amount   int
}

type Attack struct {
	Duration int
	Amount   int
}

type Draw struct {
	Amount int
}

type Discard struct {
	Amount int
}

func (a Poison) Do(ac *models.GameContext) {
	ac.Destination.Curses = append(ac.Destination.Curses, models.Effect{
		Field:    models.Health,
		Duration: a.Duration,
		Amount:   a.Amount,
	})
}

func (a Heal) Do(ac *models.GameContext) {
	if randCritical(ac.Source) {
		a.Amount *= ac.Source.Parameters[models.CriticalDamage] / 100
	}

	if a.Duration <= 1 {
		ac.Source.Parameters[models.Health] += a.Amount
		return
	}

	ac.Source.Parameters[models.Health] += a.Amount
	ac.Source.Buffs = append(ac.Source.Buffs, models.Effect{
		Field:    models.Health,
		Duration: a.Duration - 1,
		Amount:   a.Amount,
	})
}

func (a Attack) Do(ac *models.GameContext) {
	if randDodge(ac.Destination) {
		return
	}

	if randCritical(ac.Source) {
		a.Amount += a.Amount * ac.Source.Parameters[models.CriticalDamage] / 100
	}

	if a.Duration <= 1 {
		ac.Destination.Parameters[models.Health] -= a.Amount
		return
	}

	ac.Destination.Parameters[models.Health] -= a.Amount
	ac.Destination.Curses = append(ac.Destination.Curses, models.Effect{
		Field:    models.Health,
		Duration: a.Duration - 1,
		Amount:   a.Amount,
	})
}

func (a Draw) Do(ac *models.GameContext) {
	ac.Source.Hand = append(ac.Source.Hand, ac.Source.Deck[:a.Amount]...)
	// delete drawn cards from deck
	ac.Source.Deck = ac.Source.Deck[a.Amount:]
}

func (a Discard) Do(ac *models.GameContext) {
	ac.Destination.Hand = ac.Destination.Hand[a.Amount:]
}

func randCritical(c *models.Character) bool {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(100) <= c.Parameters[models.Critical]
}

func randDodge(c *models.Character) bool {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(100) <= c.Parameters[models.Dodge]
}
