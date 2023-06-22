package game

var cards = []Card{
	{
		ID:          1,
		Name:        "card1",
		Description: "card1",
		Cost:        2,
		Rarity:      Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 5},
		},
	},
	{
		ID:          2,
		Name:        "card2",
		Description: "card2",
		Cost:        3,
		Rarity:      Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          3,
		Name:        "card3",
		Description: "card3",
		Cost:        2,
		Rarity:      Common,
		Actions: []IAction{
			Heal{Duration: 1, Amount: 5},
		}, 
	},
	{
		ID:          4,
		Name:        "card4",
		Description: "card4",
		Cost:        3,
		Rarity:      Common,
		Actions: []IAction{
			Heal{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          5,
		Name:        "card5",
		Description: "card5",
		Cost:        3,
		Rarity:      Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 5},
			Draw {1},
		},
	},
	{
		ID: 		6,
		Name: 		"card6",
		Description:"card6",
		Cost:		4,
		Rarity: 	Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 8},
			Draw {1},
		},
	},
	{
		ID: 		7,
		Name: 		"card7",
		Description:"card7",
		Cost:		4,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{ 3, 4},
		},
	},
	{
		ID: 		8,
		Name: 		"card8",
		Description:"card8",
		Cost:		5,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{3, 3},
			Draw {1},
		},
	},
}