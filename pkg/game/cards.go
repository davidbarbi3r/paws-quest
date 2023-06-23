package game

var commonCards = []Card{
	{
		ID:          1,
		Name:        "Paw Swipe",
		Description: "A quick swipe",
		Cost:        2,
		Rarity:      Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 5},
		},
	},
	{
		ID:          2,
		Name:        "Sneaky Scratch",
		Description: "Quick and elusive",
		Cost:        3,
		Rarity:      Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          3,
		Name:        "CatNap",
		Description: "Just a little nap",
		Cost:        2,
		Rarity:      Common,
		Actions: []IAction{
			Heal{Duration: 1, Amount: 5},
		}, 
	},
	{
		ID:          4,
		Name:        "Nap time",
		Description: "A well earned rest",
		Cost:        3,
		Rarity:      Common,
		Actions: []IAction{
			Heal{Duration: 1, Amount: 8},
		},
	},
	{
		ID:          5,
		Name:        "Lets change the future",
		Description: "A card for the tikiton",
		Cost:        0,
		Rarity:      Common,
		Actions: []IAction{
			Draw {1},
		},
	},
	{
		ID: 		6,
		Name: 		"Smart fist",
		Description:"This fist is smart",
		Cost:		4,
		Rarity: 	Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 8},
			Draw {1},
		},
	},
	{
		ID: 		11,
		Name: 		"Triple claw strike",
		Description:"3 times better than one",
		Cost:		3,
		Rarity: 	Common,
		Actions: []IAction{
			Attack{Duration: 1, Amount: 2},
			Attack{Duration: 1, Amount: 2},
			Attack{Duration: 1, Amount: 2},
		},
	},
	{
		ID: 		13,
		Name: 		"Litter box",
		Description:"You don't want to go there",
		Cost:		2,
		Rarity: 	Common,
		Actions: []IAction{
			Poison{Duration: 3, Amount: 2},
		},
	},
}

var uncommonCards = []Card{
	{
		ID: 		7,
		Name: 		"Venomous bite",
		Description:"You will feel it",
		Cost:		4,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{1, 4},
			Poison{3, 3},
		},
	},
	{
		ID: 		8,
		Name: 		"Poisonous bite",
		Description:"You'll die slowly but surely",
		Cost:		5,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{1, 3},
			Poison{2, 3},
			Draw {1},
		},
	},
	{
		ID: 		9,
		Name: 		"Paw Paw",
		Description:"Texas style paw paw",
		Cost:		3,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{1, 4},
			Attack{1, 4},
		},
	},
	{
		ID: 		10,
		Name: 		"Vampirisim",
		Description:"Blood will flow",
		Cost:		3,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Attack{1, 3},
			Heal{1, 3},
		},
	},
	{
		ID: 		12,
		Name: 		"Possibilities",
		Description:"You can do anything",
		Cost:		0,
		Rarity: 	Uncommon,
		Actions: []IAction{
			Draw {2},
		},
	},
}

var rareCards = []Card{
	{
		ID: 		14,
		Name: 		"Tikiton fury",
		Description:"The tikiton is angry",
		Cost: 		0,
		Rarity: 	Rare,
		Actions: []IAction{
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
			Attack{1, 1},
		},
	},
	{
		ID: 		15,
		Name: 		"Catnip",
		Description:"You'll feel good",
		Cost:		1,
		Rarity: 	Rare,
		Actions: []IAction{
			Heal{1, 10},
			Draw {1},
		},
	},
	{
		ID: 		16,
		Name: 		"Catnip overdose",
		Description:"You'll feel too good",
		Cost:		2,
		Rarity: 	Rare,
		Actions: []IAction{
			Heal{1, 20},
		},
	},
}

