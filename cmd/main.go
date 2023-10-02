package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// db
	dbService "example/paws-quest/pkg/database"
	"example/paws-quest/pkg/game"

	_ "github.com/lib/pq"

	// internal packages
	"example/paws-quest/pkg/game/gamemap"
	"example/paws-quest/pkg/player"
)

// todo
// - [X] DONE -- add the logic to have combined actions (ex: attack + poison)
// - [X] DONE -- implement the logic when the player play a card or when an enemy do an action for missing attacks, critical hit...
// - [X] DONE -- add the buff logic
// - [X] DONE -- add the simple deck creation logic
// - [X] DONE -- reorganize the code to have a models readonly package
// - [ ] add the deck creation logic at the start of the game after the player has chosen their character
// - [ ] add the object model
// - [ ] add the shop logic and implement croquinette economy
// - [ ] add the rest node logic
// - [ ] implement the logic between two fights when an enemy is dead
// - [ ] implement the possibility to have multiple enemies in a same node
// - [ ] add the enemy map population

func main() {
	gamemapService := gamemap.ServiceImpl{}
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
