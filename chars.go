package main

func PlayerChar() Char {
	MainPlayer := Char{
		Information: Information{
			Name:        "John Doe",
			Age:         27,
			Description: "im an old fella", // add base damage
		},
		Stats: Stats{
			Health:       100,
			Mana:         20,
			Agility:      4,
			Strength:     8,
			Intelligence: 6,
			BaseDefence:  2,
			Defence:      0,
		},
		Storage: Storage{
			BackPack: map[string]Item{},
			Pockets:  map[string]int{},
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}
	return MainPlayer
}

func NPC() map[string]Char {
	NpcList := make(map[string]Char)

	TestNPC := Char{
		Information: Information{
			Name:        "jimme the sketch",
			Age:         50,
			Description: "He do be an old sketchy fella",
		},
		Stats: Stats{
			Health:       40,
			Mana:         30,
			Agility:      6,
			Strength:     4,
			Intelligence: 4,
		},
		Storage: Storage{
			BackPack: map[string]Item{},
			Pockets:  map[string]int{},
		},
		Location: Location{
			XAxis: 5,
			YAxis: 5,
		},
		ObjectState: ObjectState{
			Death: map[string]Char{},
		},
	}

	TargetDummy := Char{
		Information: Information{
			Name:        "Dummy",
			Age:         100,
			Description: "A target dummy made for testing",
		},
		Stats: Stats{
			Health:       100,
			Mana:         100,
			Agility:      10,
			Strength:     10,
			Intelligence: 10,
		},
		Storage: Storage{
			BackPack:  map[string]Item{},
			Pockets:   map[string]int{},
			LootTable: map[string]Item{},
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
		ObjectState: ObjectState{
			Death: map[string]Char{},
		},
	}

	NpcList["testNPC"] = TestNPC
	NpcList["TargetDummy"] = TargetDummy
	// fmt.Println(NpcList["testNPC"])

	return NpcList
}

func ItemArrays() []Item {

	ItemArray := []Item{}

	BegginerMace := Item{
		Kind: ItemKindWeapon,
		Information: Information{
			Name:        "TestMace",
			Age:         100,
			Description: "To be written",
		},
		Stats: Stats{
			Health:       5,
			Mana:         5,
			Agility:      5,
			Strength:     5,
			Intelligence: 5,
			BaseDmg:      5,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	BegginerAxe := Item{
		Kind: ItemKindWeapon,
		Information: Information{
			Name:        "TestAxe",
			Age:         100,
			Description: "To be written",
		},
		Stats: Stats{
			Health:       5,
			Mana:         5,
			Agility:      5,
			Strength:     5,
			Intelligence: 5,
			BaseDmg:      5,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}
	BegginerSword := Item{
		Kind: ItemKindWeapon,
		Information: Information{
			Name:        "TestAxe",
			Age:         100,
			Description: "To be written",
		},
		Stats: Stats{
			Health:       5,
			Mana:         5,
			Agility:      5,
			Strength:     5,
			Intelligence: 5,
			BaseDmg:      5,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	HealthPotion := Item{
		Kind: ItemKindPotion,
		Information: Information{
			Name:        "Health Potion",
			Age:         1,
			Description: "Using this potion will replenish your health points",
		},
		Stats: Stats{
			Health: 25,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	ItemArray = append(ItemArray, BegginerMace)
	ItemArray = append(ItemArray, BegginerAxe)
	ItemArray = append(ItemArray, BegginerSword)
	ItemArray = append(ItemArray, HealthPotion)

	return ItemArray
}

// func Items() map[string]Item {
// 	return map[string]Item{}
// }

func Items() map[string]Item {
	ItemList := make(map[string]Item)

	TestSword := Item{
		Information: Information{
			Name:        "Test",
			Age:         100,
			Description: "Something",
		},
		Stats: Stats{
			Health:       10,
			Mana:         10,
			Agility:      10,
			Strength:     10,
			Intelligence: 10,
			BaseDmg:      10,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}
	woodenMallet := Item{
		Information: Information{
			Name:        "Test",
			Age:         100,
			Description: "Something",
		},
		Stats: Stats{
			Health:       10,
			Mana:         10,
			Agility:      10,
			Strength:     10,
			Intelligence: 10,
			BaseDmg:      10,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}
	testHammer := Item{
		Information: Information{
			Name:        "testHammer",
			Age:         100,
			Description: "A hammer made for testing",
		},
		Stats: Stats{
			Health:       10,
			Mana:         10,
			Agility:      10,
			Strength:     10,
			Intelligence: 10,
			BaseDmg:      10,
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	// testBoow := Item{
	// 	InformaInformation: Information{
	// 		Name:
	// 		Age:
	// 		Description:
	// 	},

	// }

	ItemList["TestSword"] = TestSword
	ItemList["WoodenMallet"] = woodenMallet
	ItemList["testHammer"] = testHammer

	return ItemList
}
