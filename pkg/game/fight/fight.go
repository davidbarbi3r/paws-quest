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

func applyBuffs (c *game.Character) {
	if len(c.Buffs) == 0 {
		return
	}
	for i := 0; i < len(c.Buffs); i++ {
		if c.Buffs[i].Duration > 0 {
			if c.Buffs[i].Field == game.Health {
				c.Parameters[c.Buffs[i].Field] += c.Buffs[i].Amount
				c.Buffs[i].Duration--
			} else {
				c.Parameters[c.Buffs[i].Field] += c.Buffs[i].Amount
			}		
		}
	}
}

func applyDebuffs (c *game.Character) {
	if len(c.Buffs) == 0 {
		return
	}
	for i := 0; i < len(c.Buffs); i++ {
		
		if c.Buffs[i].Duration > 0 {
			if c.Buffs[i].Field == game.Health {
				continue 
			} 
				
			c.Parameters[c.Buffs[i].Field] -= c.Buffs[i].Amount
			c.Buffs[i].Duration--
		}
	}
}

func removeExpiredCursesAndBuffs (c *game.Character) {
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

func StartPlayerTurn (gc *game.GameContext, g *game.Game) {
	applyCurses(gc.Source)
	applyBuffs(gc.Source)
	removeExpiredCursesAndBuffs(gc.Source)

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
	
	for _, action := range card.Actions {
		action.Do(gc)
	}

	if isCharacterDead(gc.Source) {
		g.State = game.GameOver
		return
	}

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

	applyDebuffs(gc.Destination)

	// switch source and destination
	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = game.EnemyTurn
}

func EnemyTurn (gc *game.GameContext, g *game.Game) {
	applyCurses(gc.Source)
	removeExpiredCursesAndBuffs(gc.Source)

	if isCharacterDead(gc.Source) {
		g.State = game.Loot
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
		g.State = game.GameOver
		return
	}

	applyDebuffs(gc.Destination)

	gc.Source, gc.Destination = gc.Destination, gc.Source
	g.State = game.PlayerTurn
}


