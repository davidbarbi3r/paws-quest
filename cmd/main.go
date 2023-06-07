package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/joho/godotenv"

	// db
	dbService "example/paws-quest/pkg/database"
	_ "github.com/lib/pq"

	// internal packages
	player "example/paws-quest/pkg/player"
	gamemap "example/paws-quest/pkg/game/gamemap"
	game "example/paws-quest/pkg/game"
)

func main () {
	gamemapService := gamemap.GameMapServiceImpl{}
	fmt.Println(gamemapService.Create(496603))
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
	game.SetupGameRouter(router.Group("/api/game"))

	router.Run("localhost:8080")
}
