package gamemap

type NodeType string

const (
	Fight NodeType = "fight"
	Shop NodeType = "shop"
	Rest NodeType = "rest"
	Boss NodeType = "boss"
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



