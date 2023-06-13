package fight

import (
	"example/paws-quest/pkg/game/action"
	"example/paws-quest/pkg/game/card"
	"example/paws-quest/pkg/game/character"
)

func hasEnoughStamina (c *character.Character, cost int) bool {
	return c.Stamina >= cost
}

func setCharacterDead (game *action.GameContext) {
	if game.Destination.Health <= 0 {
		game.Destination.Health = 0
		game.Destination.IsDead = true
	}
	if game.Source.Health <= 0 {
		game.Source.Health = 0
		game.Source.IsDead = true
	}
}

func applyCurses (game *action.GameContext) {
	for i := 0; i < len(game.Destination.Curses); i++ {
		// game.Destination[game.Destination.Field] -= game.Destination.Curses[i].Amount
		game.Destination.Curses[i].Duration--
		if game.Destination.Curses[i].Duration <= 0 {
			// remove curse i from curses
			game.Destination.Curses = append(game.Destination.Curses[:i], game.Destination.Curses[i+1:]...)
		}
	}
}

func playCard (game *action.GameContext, card *card.Card) {
	if !hasEnoughStamina(game.Source, card.Cost) {
		return
	}
	game.Source.Stamina -= card.Cost
	
} 