package player

import (
	"github.com/gin-gonic/gin"
)

func SetupPlayersRouter(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "users",
		})
	})
}
