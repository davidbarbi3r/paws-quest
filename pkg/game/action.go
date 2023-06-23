package game

import (
	"math/rand"
	"time"
)

type Effect struct {
	Field Field
	Duration int
	Amount int
}

type GameContext struct {
	Source      *Character
	Destination *Character
}

type IAction interface {
	Do(*GameContext)
}

type Poison struct {
	Duration int
	Amount int
}

func (a Poison) Do(ac *GameContext) {
	ac.Destination.Curses = append(ac.Destination.Curses, Effect{
		Field:    Health,
		Duration: a.Duration,
		Amount:   a.Amount,
	})
}

type Heal struct {
	Duration int
	Amount int
}

func (a Heal) Do(ac *GameContext) {
	if randCritical(ac.Source) {
		a.Amount *= ac.Source.Parameters[CriticalDamage]/100
	}

	if a.Duration <= 1 {
		ac.Source.Parameters[Health] += a.Amount
		return
	}
	
	ac.Source.Parameters[Health] += a.Amount
	ac.Source.Buffs = append(ac.Source.Buffs, Effect{
		Field:    Health,
		Duration: a.Duration - 1,
		Amount:   a.Amount,
	})
}

type Attack struct {
	Duration int
	Amount int
}

func (a Attack) Do(ac *GameContext) {	
	if randDodge(ac.Destination) {
		return
	}

	if randCritical(ac.Source) {
		a.Amount *= ac.Source.Parameters[CriticalDamage]/100
	}

	if a.Duration <= 1 {
		ac.Destination.Parameters[Health] -= a.Amount
		return
	}

	ac.Destination.Parameters[Health] -= a.Amount
	ac.Destination.Curses = append(ac.Destination.Curses, Effect{
		Field:    Health,
		Duration: a.Duration - 1,
		Amount:   a.Amount,
	})
}

type Draw struct {
	Amount int
}

func (a Draw) Do(ac *GameContext) {
	ac.Source.Hand = append(ac.Source.Hand, ac.Source.Deck[:a.Amount]...)
	// delete drawn cards from deck
	ac.Source.Deck = ac.Source.Deck[a.Amount:]
}

type Discard struct {
	Amount int
}

func (a Discard) Do(ac *GameContext) {
	ac.Destination.Hand = ac.Destination.Hand[a.Amount:]
}

func randCritical(c *Character) bool {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(100) <= c.Parameters[Critical]
} 

func randDodge(c *Character) bool {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(100) <= c.Parameters[Dodge]
}

