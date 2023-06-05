package game

import (
	"example/paws-quest/pkg/player"
)

type Game interface {
	Start (player.Player) (string, error)
	Stop (player.Player) (string, error)
	Save (player.Player) (string, error)
}


