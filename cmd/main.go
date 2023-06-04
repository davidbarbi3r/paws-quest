package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/joho/godotenv"
	player "example/paws-quest/pkg/player"
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

	player.SetupPlayersRouter(router.Group("/api/player"))

	router.Run("localhost:8080")
}
