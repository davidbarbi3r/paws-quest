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
	if a.Duration > 1 {
		ac.Source.Parameters[Health] += a.Amount
		ac.Source.Curses = append(ac.Source.Curses, Effect{
			Field:    Health,
			Duration: a.Duration - 1,
			Amount:   a.Amount,
		})
	} else {
		ac.Source.Parameters[Health] += a.Amount
	}
}

type Attack struct {
	Duration int
	Amount int
}

func (a Attack) Do(ac *GameContext) {
	if a.Duration > 1 {
		ac.Destination.Parameters[Health] -= a.Amount
		ac.Destination.Curses = append(ac.Destination.Curses, Effect{
			Field:    Health,
			Duration: a.Duration - 1,
			Amount:   a.Amount,
		})
	} else {
		ac.Destination.Parameters[Health] -= a.Amount
	}
}

type Draw struct {
	Amount int
}

func (a Draw) Do(ac *GameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[a.Amount])
	ac.Source.Deck = append(ac.Source.Deck[:a.Amount], ac.Source.Deck[a.Amount+1:]...)
}

type Discard struct {
	Card int
}

func (a Discard) Do(ac *GameContext) {
	ac.Destination.Hand = append(ac.Destination.Hand[:a.Card], ac.Destination.Hand[a.Card+1:]...)
}

type AttackDraw struct {
	Dmg int
	Card int
}

func (a AttackDraw) Do(ac *GameContext) {
	ac.Destination.Parameters[Health] -= a.Dmg
	for i := 0; i < a.Card; i++ {
		ac.Destination.Hand = append(ac.Destination.Hand, ac.Destination.Deck[i])
	}
}