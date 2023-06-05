package gamemap

type NodeType string

const (
	Fight NodeType = "fight"
	Shop NodeType = "shop"
	Rest NodeType = "rest"
	Mystery NodeType = "mystery"
)

type Node struct {
	ID int
	Type NodeType
	Connections []int
}

type GameMap struct {
	Nodes map[int]Node
}

type GameMapService interface {
	Create(seed int) GameMap
	AddNode(gamemap* GameMap, node Node) GameMap
	ListNodes(gamemap GameMap) []Node
	GetNode(gamemap GameMap, id int) Node
}



