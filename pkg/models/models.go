package models

/* -------- GAME -------- */

type GameState string

const (
	CatSelection GameState = "cat-selection"
	Fight        GameState = "fight"
	EnemyTurn    GameState = "enemy-turn"
	PlayerTurn   GameState = "player-turn"
	Loot         GameState = "loot"
	Shop         GameState = "shop"
	Rest         GameState = "rest"
	Boss         GameState = "boss"
	GameOver     GameState = "game-over"
	GameWon      GameState = "game-won"
)

type GameService interface {
	Create(Player) (Game, error)
	Stop(Player) (string, error)
	Save(Player) (string, error)
}

type Game struct {
	Map         GameMap
	CatChosen   *int
	GameContext *GameContext
	State       GameState
	CatChoice   []int
	CurrentNode Node
	ID          int
	Player      int
}

/* -------- MAP -------- */

type NodeType string

const (
	FightNode NodeType = "fight"
	ShopNode  NodeType = "shop"
	RestNode  NodeType = "rest"
	BossNode  NodeType = "boss"
	// Mystery NodeType = "mystery"
)

type Node struct {
	Type        NodeType
	Connections []Node
	ID          int
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
	Health         Field = iota // from 10 to 100
	Stamina                     // from 0 to 10
	Speed                       // from 0 to 10
	Critical                    // from 0 to 100
	CriticalDamage              // from 0 to 100
	Dodge                       // from 0 to 100
	HandSize                    // from 0 to 10
)

type Character struct {
	Parameters  map[Field]int
	Name        string
	Deck        []Card
	Hand        []Card
	CardsPatern []int
	Discard     []Card
	Curses      []Effect
	Buffs       []Effect
	ID          int
	Type        Element
	IsDead      bool
}

/* -------- CARDS -------- */

type Card struct {
	Name        string
	Description string
	Actions     []IAction
	ID          int
	Cost        int
	Rarity      Rarity
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
	Field    Field
	Duration int
	Amount   int
}

type GameContext struct {
	Source      *Character
	Destination *Character
}

/* -------- PLAYER -------- */

type Player struct {
	Name       string
	Email      string
	Password   string
	ID         int
	Level      int
	Experience int
}
