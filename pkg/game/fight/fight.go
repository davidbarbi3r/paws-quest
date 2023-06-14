package fight

import (
	game "example/paws-quest/pkg/game"
)

func hasEnoughStamina (c *game.Character, cost int) bool {
	return c.Parameters[game.Stamina] >= cost
}

func isCharacterDead (gc *game.GameContext) bool {
	return gc.Destination.Parameters[game.Health] <= 0 || gc.Source.Parameters[game.Health] <= 0
}

func applyCurses (c *game.Character) {
	for i := 0; i < len(c.Curses); i++ {

		c.Parameters[c.Curses[i].Field] -= c.Curses[i].Amount

		c.Curses[i].Duration--
		if c.Curses[i].Duration <= 0 {
			// remove curse i from curses
			c.Curses = append(c.Curses[:i], c.Curses[i+1:]...)
		}
	}
}

func PlayCard (gc *game.GameContext, card *game.Card) {
	if !hasEnoughStamina(gc.Source, card.Cost) {
		return
	}
	// remove stamina
	gc.Source.Parameters[game.Stamina] -= card.Cost
	
	// apply card action
	card.Action.Do(gc)

	// check if character is dead
	if isCharacterDead(gc) {
		return
	}

	// remove this card from hand
	for i := 0; i < len(gc.Source.Hand); i++ {
		if gc.Source.Hand[i].ID == card.ID {
			gc.Source.Hand = append(gc.Source.Hand[:i], gc.Source.Hand[i+1:]...)
		}
	} 
}

func EndPlayerTurn (gc *game.GameContext, g *game.Game) (*game.GameContext, *game.Game) {
	
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
		
		// empty discard pile
		gc.Source.Discard = []game.Card{}
		
		// shuffle deck
		gc.Source.Deck = game.ShuffleCards(gc.Source.Deck)
		
		// remove drawn cards from deck
		gc.Source.Deck = []game.Card{}
	}
	gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.Parameters[game.HandSize]]...)
	// remove drawn cards from deck
	gc.Source.Deck = gc.Source.Deck[gc.Source.Parameters[game.HandSize]:]

	// set g state to enemy turn
	g.State = game.EnemyTurn

	return gc, g
}

func EndEnemyTurn (gc *game.GameContext, g *game.Game) (*game.GameContext, *game.Game) {
	// apply Destination curses
	applyCurses(gc.Destination)

	return gc, g
}


