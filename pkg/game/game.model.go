package game

import (
	"example/paws-quest/pkg/player"
	"example/paws-quest/pkg/game/gamemap"
)

type GameState string

const (
	CatSelection GameState = "cat-selection"
	Fight GameState = "fight"
	Shop GameState = "shop"
	Rest GameState = "rest"
	Boss GameState = "boss"
	GameOver GameState = "game-over"
	GameWon GameState = "game-won"
)

type GameService interface {
	Create (player.Player) (Game, error)
	Stop (player.Player) (string, error)
	Save (player.Player) (string, error)
}

type Game struct {
	ID int
	Player int
	CurrentNode gamemap.Node
	Map gamemap.GameMap
	State GameState
	CatChosen *int
	CatChoice []int
}



