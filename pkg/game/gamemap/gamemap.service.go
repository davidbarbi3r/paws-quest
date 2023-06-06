package gamemap

import (
	"math/rand"
)

type GameMapServiceImpl struct {}

func (gms *GameMapServiceImpl) Create (seed int64) GameMap {
	NewRand := rand.New(rand.NewSource(seed))
	nodeNumber := NewRand.Intn(5) + 5

	nodes := make(map[int]Node)

	for i := 0; i < nodeNumber; i++ {
		node := Node{
			ID: i,
			Connections: make([]Node, 0),
		}

		typeRand := NewRand.Intn(10)
		if typeRand >= 0 && typeRand <= 5 {
			node.Type = Fight
		} else if typeRand == 6 || typeRand == 7 {
			node.Type = Shop
		} else {
			node.Type = Rest
		}

		nodes[i] = node
	}

	for i := 0; i < nodeNumber; i++ {
		node := nodes[i]
		connections := NewRand.Intn(2) + 1
		for j := 0; j < connections; j++ {
			connection := nodes[NewRand.Intn(nodeNumber)]
			node.Connections = append(node.Connections, connection)
		}
		nodes[i] = node
	}

	return GameMap{
		Nodes: nodes,
	}
}

func ListNodes(gamemap GameMap) []Node {
	return []Node{}
}

func GetNode(gamemap GameMap, id int) Node {
	return Node{}
}

func AddNode(gamemap* GameMap, node Node) GameMap {
	return GameMap{}
}


