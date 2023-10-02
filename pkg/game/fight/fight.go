package fight

import (
	"example/paws-quest/pkg/card"
	"example/paws-quest/pkg/models"
)

func hasEnoughStamina(c *models.Character, cost int) bool {
	return c.Parameters[models.Stamina] >= cost
}

func isCharacterDead(c *models.Character) bool {
	return c.Parameters[models.Health] <= 0
}

func applyCurses(c *models.Character) {
	if len(c.Curses) == 0 {
		return
	}
	for i := 0; i < len(c.Curses); i++ {

		if c.Curses[i].Duration > 0 {
			c.Parameters[c.Curses[i].Field] -= c.Curses[i].Amount
			c.Curses[i].Duration--
		}
	}
}

func applyBuffs(c *models.Character) {
	if len(c.Buffs) == 0 {
		return
	}
	for i := 0; i < len(c.Buffs); i++ {
		if c.Buffs[i].Duration > 0 {
			if c.Buffs[i].Field == models.Health {
				c.Parameters[c.Buffs[i].Field] += c.Buffs[i].Amount
				c.Buffs[i].Duration--
			} else {
				c.Parameters[c.Buffs[i].Field] += c.Buffs[i].Amount
			}
		}
	}
}

func applyDebuffs(c *models.Character) {
	if len(c.Buffs) == 0 {
		return
	}
	for i := 0; i < len(c.Buffs); i++ {

		if c.Buffs[i].Duration > 0 {
			if c.Buffs[i].Field == models.Health {
				continue
			}

			c.Parameters[c.Buffs[i].Field] -= c.Buffs[i].Amount
			c.Buffs[i].Duration--
		}
	}
}

func removeExpiredCursesAndBuffs(c *models.Character) {
	if len(c.Curses) != 0 {
		for i := 0; i < len(c.Curses); i++ {
			if c.Curses[i].Duration == 0 {
				c.Curses = append(c.Curses[:i], c.Curses[i+1:]...)
			}
		}
	}
	if len(c.Buffs) != 0 {
		for i := 0; i < len(c.Buffs); i++ {
			if c.Buffs[i].Duration == 0 {
				c.Buffs = append(c.Buffs[:i], c.Buffs[i+1:]...)
			}
		}
	}
}

func StartPlayerTurn(gc *models.GameContext, g *models.Game) {
	applyCurses(gc.Source)
	applyBuffs(gc.Source)
	removeExpiredCursesAndBuffs(gc.Source)

	if isCharacterDead(gc.Source) {
		g.State = models.GameOver
		return
	}

	g.State = models.PlayerTurn
}

func PlayCard(gc *models.GameContext, g *models.Game, card models.Card) {
	if !hasEnoughStamina(gc.Source, card.Cost) {
		return
	}

	gc.Source.Parameters[models.Stamina] -= card.Cost

	for _, action := range card.Actions {
		action.Do(gc)
	}

	if isCharacterDead(gc.Source) {
		g.State = models.GameOver
		return
	}

	for i := 0; i < len(gc.Source.Hand); i++ {
		if gc.Source.Hand[i].ID == card.ID {
			gc.Source.Hand = append(gc.Source.Hand[:i], gc.Source.Hand[i+1:]...)
		}
	}
}

func EndPlayerTurn(gc *models.GameContext, g *models.Game) {
	// reset stamina // todo setup initial properties when needed
	gc.Source.Parameters[models.Stamina] = gc.Source.Parameters[models.Speed]

	// put unused cards in hand in discard pile
	gc.Source.Discard = append(gc.Source.Discard, gc.Source.Hand...)

	// remove cards from hand
	gc.Source.Hand = []models.Card{}

	// draw cards
	// do I have enough cards in deck to draw? Yes
	if len(gc.Source.Deck) > gc.Source.Parameters[models.HandSize] {
		gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.Parameters[models.HandSize]]...)
		// remove drawn cards from deck
		gc.Source.Deck = gc.Source.Deck[gc.Source.Parameters[models.HandSize]:]

		// do I have enough cards in deck to draw? No
	} else {
		// put discard pile in deck
		gc.Source.Deck = append(gc.Source.Deck, gc.Source.Discard...)

		gc.Source.Discard = []models.Card{}

		gc.Source.Deck = card.Shuffle(gc.Source.Deck)

		// remove drawn cards from deck
		gc.Source.Deck = []models.Card{}
	}
	gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.Parameters[models.HandSize]]...)
	// remove drawn cards from deck
	gc.Source.Deck = gc.Source.Deck[gc.Source.Parameters[models.HandSize]:]

	applyDebuffs(gc.Destination)

	// switch source and destination
	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = models.EnemyTurn
}

func EnemyTurn(gc *models.GameContext, g *models.Game) {
	applyCurses(gc.Source)
	removeExpiredCursesAndBuffs(gc.Source)

	if isCharacterDead(gc.Source) {
		g.State = models.Loot
		return
	}

	if len(gc.Source.CardsPatern) != 0 {
		card := gc.Source.Hand[gc.Source.CardsPatern[0]]

		for _, action := range card.Actions {
			action.Do(gc)
		}
		// put card played at the end of the pattern
		gc.Destination.CardsPatern = append(gc.Destination.CardsPatern[1:], gc.Destination.CardsPatern[0])
	}

	if isCharacterDead(gc.Destination) {
		g.State = models.GameOver
		return
	}

	applyDebuffs(gc.Destination)

	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = models.PlayerTurn
}
