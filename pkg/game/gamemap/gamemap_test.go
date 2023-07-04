package gamemap

import (
	"example/paws-quest/pkg/models"
	"fmt"
	"testing"
)

func TestCreateGameMap(t *testing.T) {
	seed := int64(45582)

	gms := &GameMapServiceImpl{}

	gameMap := gms.Create(seed)

	// Verify node number
	nodeNumber := len(gameMap.Nodes)
	if nodeNumber < 5 || nodeNumber > 10 {
		t.Errorf("Incorrect node number. Node number: %d, Expected: between 5 & 10", nodeNumber)
	}

	// Verify connections
	for i := 0; i < nodeNumber; i++ {
		node := gameMap.Nodes[i]

		// Verify that the node is connected to a superior node
		for _, connection := range node.Connections {
			if connection.ID <= i && connection.ID != 0 {
				t.Errorf("Node %d connected to inferior or equal Node (%d)", node.ID, connection.ID)
			}
		}

		// Last nodes have to be connected to the boss node
		if i == nodeNumber - 1 {
			connectedToBoss := false
			for _, connection := range node.Connections {
				if connection.Type == models.BossNode {
					connectedToBoss = true
				}
			}
			if !connectedToBoss {
				t.Errorf("Last node is not connected to the boss node")
			}
		}

		// Last node have to be of type Boss
		if i == nodeNumber && node.Type != models.BossNode {
			t.Errorf("The last Node is not of type Boss. Type: %s, Expected: Boss", node.Type)
		}
	}

	// two game maps with the same seed should be the same
	gameMap2 := gms.Create(seed)
	if gameMap.Nodes[0].ID != gameMap2.Nodes[0].ID {
		t.Errorf("Two game maps with the same seed are not the same")
	}

	fmt.Println("TestCreateGameMap ✅.")
}