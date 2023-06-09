package card

import (
	"example/paws-quest/pkg/game/character"
)

type player struct {
	Hp   int
	Mana int
	Curse []curse
	Deck []Card
	Hand []Card
}

type curse struct {
	Field Field
	Duration int
}

type gameContext struct {
	Source      *character.Character
	Destination *character.Character
}

type leechLife struct {
	Hp int
}

func (a *leechLife) Do(ac *gameContext) {
	ac.Source.Hp += a.Hp
	ac.Destination.Hp -= a.Hp
}

type poison struct {
	Dmg int
	Dot int
	Duration int
}

func (a *poison) Do(ac *gameContext) {
	ac.Destination.Hp -= a.Dmg
	ac.Destination.Curse = append(ac.Destination.Curse, curse{
		Field:    Health,
		Duration: a.Duration,
	})
}

type heal struct {
	Hp int
}

func (a *heal) Do(ac *gameContext) {
	ac.Destination.Hp += a.Hp
}

type attack struct {
	Dmg int
}

func (a *attack) Do(ac *gameContext) {
	ac.Destination.Hp -= a.Dmg
}

type draw struct {
	Card int
}

func (a *draw) Do(ac *gameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[a.Card])
}

type discard struct {
	Card int
}

func (a *discard) Do(ac *gameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand[:a.Card], ac.Destination.Hand[a.Card+1:]...)
}

type attackDraw struct {
	Dmg int
	Card int
}

func (a *attackDraw) Do(ac *gameContext) {
	ac.Destination.Hp -= a.Dmg
	for i := 0; i < a.Card; i++ {
		ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[i])
	}
}