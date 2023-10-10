package game

import (
	"net/http"

	// internal packages
	"example/paws-quest/pkg/game/gamemap"
	"example/paws-quest/pkg/models"
	// external packages (third party)
	"github.com/gin-gonic/gin"
)

type ServiceImpl struct{}

func (gs *ServiceImpl) Create(ctx *gin.Context) {
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
	gms := gamemap.ServiceImpl{}
	newMap := gms.Create(496603)

	// Create a new Game (next with the GameService)
	game := models.Game{
		ID:          423,
		Player:      playerId,
		CurrentNode: newMap.Nodes[0],
		Map:         newMap,
		State:       models.CatSelection,
		CatChosen:   nil,
		CatChoice:   []int{1, 2, 3},
	}

	// Save the game to the database using the GameService

	ctx.JSON(http.StatusOK, game)
}
