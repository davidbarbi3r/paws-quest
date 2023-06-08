package game

import (
	// internal packages
	"example/paws-quest/pkg/game/gamemap"

	// external packages (standard library)
	"net/http"

	// external packages (third party)
	"github.com/gin-gonic/gin"
)

type GameServiceImpl struct {}

func (gs *GameServiceImpl) Create (ctx *gin.Context) {
	playerId := ctx.GetInt("playerId")
	
	// Get the player from the database using the PlayerService
	// TODO: Implement PlayerService
	// ps := player.PlayerServiceImpl{}
	// player, err := ps.GetByID(playerId)
	// if err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"message": "Player not found",
	// 	})
	// 	return
	// }

	// Generate a new GameMap using the GameMapService
	gms := gamemap.GameMapServiceImpl{}
	newMap := gms.Create(496603)

	// Create a new Game (next with the GameService)
	game := Game{
		ID:          423, 
		Player:      playerId,
		CurrentNode: newMap.Nodes[0], 
		Map:         newMap, 
		State:       CatSelection,
		CatChosen:   nil, 
		CatChoice:   []int{1, 2, 3},
	}

	// Save the game to the database using the GameService

	ctx.JSON(http.StatusOK, game)
}
