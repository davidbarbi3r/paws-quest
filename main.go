package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	user "example/paws-quest/user"
)

func main () {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	user.SetupUsersRouter(router.Group("/api/user"))

	router.Run("localhost:8080")
}
