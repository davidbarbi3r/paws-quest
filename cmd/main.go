package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/joho/godotenv"
	user "example/paws-quest/pkg/user"
	dbService "example/paws-quest/pkg/database"
	_ "github.com/lib/pq"
)

func main () {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	connectionString := os.Getenv("POSTGRES_URL")
	
	router := gin.Default()
	
	database := dbService.NewPostgreSQLService(connectionString) 
	_, err = database.Connect()
	if err != nil {
		panic(err)
	} 
	fmt.Println("[Database] Connected to database")
	
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	user.SetupUsersRouter(router.Group("/api/user"))

	router.Run("localhost:8080")
}
