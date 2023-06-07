package game

import (
	"github.com/gin-gonic/gin"
)

func SetupGameRouter (router *gin.RouterGroup) {
	gameService := GameServiceImpl{}
	router.POST("/create", gameService.Create)
}
