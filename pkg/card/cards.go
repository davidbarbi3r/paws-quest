package card

import (
	"example/paws-quest/pkg/models"
	"example/paws-quest/pkg/action"
)

var commonCards = []models.Card{
	{
		ID:          1,
		Name:        "Paw Swipe",
		Description: "A quick swipe",
		Cost:        2,
		Rarity:      models.Common,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 5},
		},
	},
	{
		ID:          2,
		Name:        "Sneaky Scratch",
		Description: "Quick and elusive",
		Cost:        3,
		Rarity:      models.Common,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          3,
		Name:        "CatNap",
		Description: "Just a little nap",
		Cost:        2,
		Rarity:      models.Common,
		Actions: []models.IAction{
			action.Heal{Duration: 1, Amount: 5},
		}, 
	},
	{
		ID:          4,
		Name:        "Nap time",
		Description: "A well earned rest",
		Cost:        3,
		Rarity:      models.Common,
		Actions: []models.IAction{
			action.Heal{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          5,
		Name:        "Lets change the future",
		Description: "A card for the tikiton",
		Cost:        0,
		Rarity:      models.Common,
		Actions: []models.IAction{
			action.Draw {Amount: 1},
		},
	},
	{
		ID: 		6,
		Name: 		"Smart fist",
		Description:"This fist is smart",
		Cost:		4,
		Rarity: 	models.Common,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 8},
			action.Draw {Amount: 1},
		},
	},
	{
		ID: 		11,
		Name: 		"Triple claw strike",
		Description:"3 times better than one",
		Cost:		3,
		Rarity: 	models.Common,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 2},
			action.Attack{Duration: 1, Amount: 2},
			action.Attack{Duration: 1, Amount: 2},
		},
	},
	{
		ID: 		13,
		Name: 		"Litter box",
		Description:"You don't want to go there",
		Cost:		2,
		Rarity: 	models.Common,
		Actions: []models.IAction{
			action.Poison{Duration: 3, Amount: 2},
		},
	},
}

var uncommonCards = []models.Card{
	{
		ID: 		7,
		Name: 		"Venomous bite",
		Description:"You will feel it",
		Cost:		4,
		Rarity: 	models.Uncommon,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 4},
			action.Poison{Duration: 3, Amount: 3},
		},
	},
	{
		ID: 		8,
		Name: 		"Poisonous bite",
		Description:"You'll die slowly but surely",
		Cost:		5,
		Rarity: 	models.Uncommon,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 3},
			action.Poison{Duration: 2, Amount: 3},
			action.Draw {Amount: 1},
		},
	},
	{
		ID: 		9,
		Name: 		"Paw Paw",
		Description:"Texas style paw paw",
		Cost:		3,
		Rarity: 	models.Uncommon,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 4},
			action.Attack{Duration: 1, Amount: 4},
		},
	},
	{
		ID: 		10,
		Name: 		"Vampirisim",
		Description:"Blood will flow",
		Cost:		3,
		Rarity: 	models.Uncommon,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 3},
			action.Heal{Duration: 1, Amount: 3},
		},
	},
	{
		ID: 		12,
		Name: 		"Possibilities",
		Description:"You can do anything",
		Cost:		0,
		Rarity: 	models.Uncommon,
		Actions: []models.IAction{
			action.Draw {Amount: 2},
		},
	},
}

var rareCards = []models.Card{
	{
		ID: 		14,
		Name: 		"Tikiton fury",
		Description:"The tikiton is angry",
		Cost: 		0,
		Rarity: 	models.Rare,
		Actions: []models.IAction{
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
			action.Attack{Duration: 1, Amount: 1},
		},
	},
	{
		ID: 		15,
		Name: 		"Catnip",
		Description:"You'll feel good",
		Cost:		1,
		Rarity: 	models.Rare,
		Actions: []models.IAction{
			action.Heal{Duration: 1, Amount: 10},
			action.Draw {Amount: 1},
		},
	},
	{
		ID: 		16,
		Name: 		"Catnip overdose",
		Description:"You'll feel too good",
		Cost:		2,
		Rarity: 	models.Rare,
		Actions: []models.IAction{
			action.Heal{Duration: 1, Amount: 20},
		},
	},
}

