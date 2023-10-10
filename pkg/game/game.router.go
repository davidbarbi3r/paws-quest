package game

import (
	"github.com/gin-gonic/gin"
)

func SetupGameRouter(router *gin.RouterGroup) {
	gameService := ServiceImpl{}
	router.POST("/create", gameService.Create)
}
