package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// db
	dbService "example/paws-quest/pkg/database"

	_ "github.com/lib/pq"

	// internal packages
	game "example/paws-quest/pkg/game"
	action "example/paws-quest/pkg/game/action"
	card "example/paws-quest/pkg/game/card"
	character "example/paws-quest/pkg/game/character"
	gamemap "example/paws-quest/pkg/game/gamemap"
	player "example/paws-quest/pkg/player"
)

func main () {
	gamemapService := gamemap.GameMapServiceImpl{}
	fmt.Println(gamemapService.Create(496603))
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}


var exampleAttackCard = card.Card {
	ID: 1,
	Name: "Paw Swipe",
	Description: "An awesome pow attack",
	Type: card.Attack,
	Cost: 1,
	Rarity: int(card.Common),
	Action: action.Attack{
		Dmg: 1,
	},
}

exampleAttackCard.Action.Do(&action.GameContext{
	Source: &character.Character{
		Health: 10,
	},
	Destination: &character.Character{
		Health: 10,
	},
})

// var exampleAttackCard2 = card.Card {
// 	ID: 2,
// 	Name: "Cat Scratch"
// 	Description: "A basic attack card",
// 	Type: card.Attack,
// 	Cost: 2,
// 	Rarity: int(card.Common),
// }

// var exampleAttackCardBetterDmg = card.Card {
// 	ID: 3,
// 	Name: "Cat bomb",
// 	Description: "A basic++ attack card",
// 	Type: card.Attack,
// 	Cost: 3,
// 	Rarity: int(card.Uncommon),
// }

// var exampleDuplicateCard = card.Card {
// 	ID: 4,
// 	Name: "Duplicate",
// 	Description: "A basic card",
// 	Type: card.Attack,
// 	Cost: 1,
// 	Rarity: int(card.Common),
// }

// var exampleAttackDotCard = card.Card {
// 	ID: 5,
// 	Name: "Cat Poisoned Scratch",
// 	Description: "A basic attack card",
// 	Type: card.Attack,
// 	Cost: 2,
// 	Rarity: int(card.Common),
// }
	
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
