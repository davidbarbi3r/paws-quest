package fight

import (
	game "example/paws-quest/pkg/game"
)

func hasEnoughStamina (c *game.Character, cost int) bool {
	return c.Parameters[game.Stamina] >= cost
}

func isCharacterDead (c *game.Character) bool {
	return c.Parameters[game.Health] <= 0
}

func applyCurses (c *game.Character) {
	for i := 0; i < len(c.Curses); i++ {

		if c.Curses[i].Duration > 0 {
			c.Parameters[c.Curses[i].Field] -= c.Curses[i].Amount
			c.Curses[i].Duration--
		}
	}
}

func removeExpiredCurses (c *game.Character) {
	for i := 0; i < len(c.Curses); i++ {
		if c.Curses[i].Duration == 0 {
			c.Curses = append(c.Curses[:i], c.Curses[i+1:]...)
		}
	}
}

func StartPlayerTurn (gc *game.GameContext, g *game.Game) {
	applyCurses(gc.Source)
	removeExpiredCurses(gc.Source)

	if isCharacterDead(gc.Source) {
		g.State = game.GameOver
		return
	}

	g.State = game.PlayerTurn
}

func PlayCard (gc *game.GameContext, g *game.Game, card game.Card) {
	if !hasEnoughStamina(gc.Source, card.Cost) {
		return
	}

	gc.Source.Parameters[game.Stamina] -= card.Cost
	
	card.Action.Do(gc)

	if isCharacterDead(gc.Source) {
		g.State = game.GameOver
		return
	}

	// remove card from hand
	for i := 0; i < len(gc.Source.Hand); i++ {
		if gc.Source.Hand[i].ID == card.ID {
			gc.Source.Hand = append(gc.Source.Hand[:i], gc.Source.Hand[i+1:]...)
		}
	} 
}

func EndPlayerTurn (gc *game.GameContext, g *game.Game) {
	// reset stamina // todo setup initial properties when needed
	gc.Source.Parameters[game.Stamina] = gc.Source.Parameters[game.Speed]
	
	// put unused cards in hand in discard pile
	gc.Source.Discard = append(gc.Source.Discard, gc.Source.Hand...)
	
	// remove cards from hand
	gc.Source.Hand = []game.Card{}
	
	// draw cards
	// do I have enough cards in deck to draw? Yes
	if len(gc.Source.Deck) > gc.Source.Parameters[game.HandSize] {
		gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.Parameters[game.HandSize]]...)
		// remove drawn cards from deck
		gc.Source.Deck = gc.Source.Deck[gc.Source.Parameters[game.HandSize]:]
	
	// do I have enough cards in deck to draw? No
	} else {
		// put discard pile in deck
		gc.Source.Deck = append(gc.Source.Deck, gc.Source.Discard...)
		
		gc.Source.Discard = []game.Card{}
		
		gc.Source.Deck = game.Shuffle(gc.Source.Deck)
		
		// remove drawn cards from deck
		gc.Source.Deck = []game.Card{}
	}
	gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.Parameters[game.HandSize]]...)
	// remove drawn cards from deck
	gc.Source.Deck = gc.Source.Deck[gc.Source.Parameters[game.HandSize]:]

	// switch source and destination
	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = game.EnemyTurn
}

func EnemyTurn (gc *game.GameContext, g *game.Game) {
	applyCurses(gc.Source)
	removeExpiredCurses(gc.Source)

	if isCharacterDead(gc.Source) {
		g.State = game.Loot
		return
	}

	card := gc.Source.Hand[gc.Destination.CardsPatern[0]]
	card.Action.Do(gc)

	if isCharacterDead(gc.Destination) {
		g.State = game.GameOver
		return
	}

	// put card played at the end of the pattern
	gc.Destination.CardsPatern = append(gc.Destination.CardsPatern[1:], gc.Destination.CardsPatern[0])

	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = game.PlayerTurn
}


