package game

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

type LeechLife struct {
	Hp int
}

func (a LeechLife) Do(ac *GameContext) {
	ac.Source.Parameters[Health] += a.Hp
	ac.Destination.Parameters[Health] -= a.Hp
}

type Poison struct {
	Dmg int
	Dot int
	Duration int
}

func (a *Poison) Do(ac *GameContext) {
	ac.Destination.Parameters[Health] -= a.Dmg
	ac.Destination.Curses = append(ac.Destination.Curses, Effect{
		Field:    Health,
		Duration: a.Duration,
		Amount:   a.Dot,
	})
}

type Heal struct {
	Hp int
}

func (a *Heal) Do(ac *GameContext) {
	ac.Destination.Parameters[Health] += a.Hp
}

type Attack struct {
	Dmg int
}

func (a Attack) Do(ac *GameContext) {
	ac.Destination.Parameters[Health] -= a.Dmg
}

// type Draw struct {
// 	Card int
// }

// func (a *Draw) Do(ac *GameContext) {
// 	ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[a.Card])
// }

// type Discard struct {
// 	Card int
// }

// func (a *Discard) Do(ac *GameContext) {
// 	ac.Destination.Hand = append(ac.Destination.Hand[:a.Card], ac.Destination.Hand[a.Card+1:]...)
// }

// type AttackDraw struct {
// 	Dmg int
// 	Card int
// }

// func (a *AttackDraw) Do(ac *GameContext) {
// 	ac.Destination.Parameters[Health] -= a.Dmg
// 	for i := 0; i < a.Card; i++ {
// 		ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[i])
// 	}
// }