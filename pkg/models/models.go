package models


/* -------- GAME -------- */

type GameState string

const (
	CatSelection GameState = "cat-selection"
	Fight GameState = "fight"
	EnemyTurn GameState = "enemy-turn"
	PlayerTurn GameState = "player-turn"
	Loot GameState = "loot"
	Shop GameState = "shop"
	Rest GameState = "rest"
	Boss GameState = "boss"
	GameOver GameState = "game-over"
	GameWon GameState = "game-won"
)

type GameService interface {
	Create (Player) (Game, error)
	Stop (Player) (string, error)
	Save (Player) (string, error)
}

type Game struct {
	ID int
	Player int
	CurrentNode Node
	Map GameMap
	State GameState
	CatChosen *int
	CatChoice []int
	GameContext *GameContext
}


/* -------- MAP -------- */

type NodeType string

const (
	FightNode NodeType = "fight"
	ShopNode NodeType = "shop"
	RestNode NodeType = "rest"
	BossNode NodeType = "boss"
	// Mystery NodeType = "mystery"
)

type Node struct {
	ID int
	Type NodeType
	Connections []Node
}

type GameMap struct {
	Nodes map[int]Node
}

type GameMapService interface {
	Create(seed int64) GameMap
}


/* -------- CHARACTERS -------- */

type Element int

const (
	Normal Element = iota
	Fire  
	Water  
	Earth  
	Air 
)

type Field int

const (
	Health Field = iota // from 10 to 100
	Stamina // from 0 to 10
	Speed // from 0 to 10
	Critical // from 0 to 100
	CriticalDamage // from 0 to 100
	Dodge // from 0 to 100
	HandSize // from 0 to 10
)

type Character struct {
	ID int 
	Name string 
	Type Element
	
	Parameters map[Field]int
	
	Deck []Card
    Hand []Card 
	CardsPatern []int
	Discard []Card

	Curses []Effect 
	Buffs []Effect 
	// Items []Item

	IsDead bool
}


/* -------- CARDS -------- */

type Card struct {
	ID int 
	Name string 
	Description string 
	Cost int
	Rarity Rarity

	Actions []IAction
}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
)


/* -------- ACTIONS -------- */

type IAction interface {
	Do(*GameContext)
}

type Effect struct {
	Field Field
	Duration int
	Amount int
}

type GameContext struct {
	Source      *Character
	Destination *Character
}


/* -------- PLAYER -------- */

type Player struct {
	ID int 
	Name string 
	Email string 
	Password string 
	Level int 
	Experience int 
}