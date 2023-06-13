package fight

import (
	game "example/paws-quest/pkg/game"
)

func hasEnoughStamina (c *game.Character, cost int) bool {
	return c.Stamina >= cost
}

func isCharacterDead (gc *game.GameContext) bool {
	return gc.Destination.Health <= 0 || gc.Source.Health <= 0
}

// func applyCurses (game *game.GameContext) {
// 	for i := 0; i < len(game.Destination.Curses); i++ {
// 		game.Destination[game.Destination.Field] -= game.Destination.Curses[i].Amount
// 		game.Destination.Curses[i].Duration--
// 		if game.Destination.Curses[i].Duration <= 0 {
// 			// remove curse i from curses
// 			game.Destination.Curses = append(game.Destination.Curses[:i], game.Destination.Curses[i+1:]...)
// 		}
// 	}
// }

func applySourceCurses (gc *game.GameContext) {
	for i := 0; i < len(gc.Source.Curses); i++ {
		switch gc.Source.Curses[i].Field {
		case "health":
			gc.Source.Health -= gc.Source.Curses[i].Amount
		case "stamina":
			gc.Source.Stamina -= gc.Source.Curses[i].Amount
		case "speed":
			gc.Source.Speed -= gc.Source.Curses[i].Amount
		case "strength":
			gc.Source.Strength -= gc.Source.Curses[i].Amount
		case "agility":
			gc.Source.Agility -= gc.Source.Curses[i].Amount
		case "intelligence":
			gc.Source.Intelligence -= gc.Source.Curses[i].Amount
		case "handSize":
			gc.Source.HandSize -= gc.Source.Curses[i].Amount
		}
		gc.Source.Curses[i].Duration--
		if gc.Source.Curses[i].Duration <= 0 {
			// remove curse i from curses
			gc.Source.Curses = append(gc.Source.Curses[:i], gc.Source.Curses[i+1:]...)
		}
	}
}

func applyDestinationCurses (gc *game.GameContext) {
	for i := 0; i < len(gc.Destination.Curses); i++ {
		switch gc.Destination.Curses[i].Field {
		case "health":
			gc.Destination.Health -= gc.Destination.Curses[i].Amount
		case "stamina":
			gc.Destination.Stamina -= gc.Destination.Curses[i].Amount
		case "speed":
			gc.Destination.Speed -= gc.Destination.Curses[i].Amount
		case "strength":
			gc.Destination.Strength -= gc.Destination.Curses[i].Amount
		case "agility":
			gc.Destination.Agility -= gc.Destination.Curses[i].Amount
		case "intelligence":
			gc.Destination.Intelligence -= gc.Destination.Curses[i].Amount
		case "handSize":
			gc.Destination.HandSize -= gc.Destination.Curses[i].Amount
		}
		gc.Destination.Curses[i].Duration--
		if gc.Destination.Curses[i].Duration <= 0 {
			// remove curse i from curses
			gc.Destination.Curses = append(gc.Destination.Curses[:i], gc.Destination.Curses[i+1:]...)
		}
	}
}

func PlayCard (gc *game.GameContext, card *game.Card) {
	if !hasEnoughStamina(gc.Source, card.Cost) {
		return
	}
	// remove stamina
	gc.Source.Stamina -= card.Cost
	
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
	
	// reset stamina
	gc.Source.Stamina = gc.Source.Speed
	
	// put unused cards in hand in discard pile
	gc.Source.Discard = append(gc.Source.Discard, gc.Source.Hand...)
	
	// remove cards from hand
	gc.Source.Hand = []game.Card{}
	
	// draw cards
	// do I have enough cards in deck to draw? Yes
	if len(gc.Source.Deck) > gc.Source.HandSize {
		gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.HandSize]...)
		// remove drawn cards from deck
		gc.Source.Deck = gc.Source.Deck[gc.Source.HandSize:]
	
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
	gc.Source.Hand = append(gc.Source.Hand, gc.Source.Deck[:gc.Source.HandSize]...)
	// remove drawn cards from deck
	gc.Source.Deck = gc.Source.Deck[gc.Source.HandSize:]

	// set g state to enemy turn
	g.State = game.EnemyTurn

	return gc, g
}

// func EndEnemyTurn (gc *game.GameContext, g *game.Game) (*game.GameContext, *game.Game) {
// 	// apply Destination curses
// 	applyDestinationCurses(gc)
// }


