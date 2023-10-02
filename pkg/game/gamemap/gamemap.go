package gamemap

import (
	"math/rand"

	"example/paws-quest/pkg/models"
)

type ServiceImpl struct{}

func (gms *ServiceImpl) Create(seed int64) models.GameMap {
	NewRand := rand.New(rand.NewSource(seed))
	nodeNumber := NewRand.Intn(5) + 5

	nodes := make(map[int]models.Node)

	for i := 0; i < nodeNumber; i++ {
		node := models.Node{
			ID:          i,
			Connections: make([]models.Node, 0),
		}

		typeRand := NewRand.Intn(10)
		if typeRand >= 0 && typeRand <= 5 {
			node.Type = models.FightNode
		} else if typeRand == 6 || typeRand == 7 {
			node.Type = models.ShopNode
		} else {
			node.Type = models.RestNode
		}

		nodes[i] = node
	}

	for i := 0; i < nodeNumber; i++ {
		node := nodes[i]
		connections := NewRand.Intn(2) + 1
		for j := 0; j < connections; j++ {
			maxDistance := 2
			newConnection := i + 1 + NewRand.Intn(maxDistance)

			if len(node.Connections) == 0 || (nodes[newConnection].ID > node.ID && node.Connections[0].ID != newConnection) {
				node.Connections = append(node.Connections, nodes[newConnection])
			}
			if i == nodeNumber-1 {
				node.Connections = append(node.Connections, models.Node{
					ID:          nodeNumber,
					Type:        models.BossNode,
					Connections: nil,
				})
			}
		}
		nodes[i] = node
	}

	return models.GameMap{
		Nodes: nodes,
	}
}
