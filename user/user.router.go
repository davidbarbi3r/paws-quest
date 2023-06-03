// export a router for the user resource:

// Path: users/user.router.go

package user

import (
	"github.com/gin-gonic/gin"
)

func SetupUsersRouter (router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "users",
		})
	})
}
