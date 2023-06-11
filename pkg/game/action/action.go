package action

import (
	"example/paws-quest/pkg/game/character"
	"example/paws-quest/pkg/game/card"
	"example/paws-quest/pkg/game"
)

type player struct {
	Hp   int
	Mana int
	Curse []character.Effect
	Deck []card.Card
	Hand []card.Card
}

type leechLife struct {
	Hp int
}

func (a *leechLife) Do(ac *game.GameContext) {
	ac.Source.Health += a.Hp
	ac.Destination.Health -= a.Hp
}

type poison struct {
	Dmg int
	Dot int
	Duration int
}

func (a *poison) Do(ac *game.GameContext) {
	ac.Destination.Health -= a.Dmg
	ac.Destination.Curses = append(ac.Destination.Curses, character.Effect{
		Field:    character.Health,
		Duration: a.Duration,
		Amount:   a.Dot,
	})
}

type heal struct {
	Hp int
}

func (a *heal) Do(ac *game.GameContext) {
	ac.Destination.Health += a.Hp
}

type attack struct {
	Dmg int
}

func (a *attack) Do(ac *game.GameContext) {
	ac.Destination.Health -= a.Dmg
}

type draw struct {
	Card int
}

func (a *draw) Do(ac *game.GameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[a.Card])
}

type discard struct {
	Card int
}

func (a *discard) Do(ac *game.GameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand[:a.Card], ac.Destination.Hand[a.Card+1:]...)
}

type attackDraw struct {
	Dmg int
	Card int
}

func (a *attackDraw) Do(ac *game.GameContext) {
	ac.Destination.Health -= a.Dmg
	for i := 0; i < a.Card; i++ {
		ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[i])
	}
}