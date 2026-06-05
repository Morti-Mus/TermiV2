package main

func PlayerChar() Char {
	MainPlayer := Char{
		Information: Information{
			Name:        "John Doe",
			Age:         27,
			Description: "im an old fella",
		},
		Stats: Stats{
			Health:       100,
			Mana:         20,
			Agility:      4,
			Strength:     8,
			Intelligence: 6,
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
			Death: map[int]int{},
		},
	}

	NpcList["testNPC"] = TestNPC
	NpcList["TargetDummy"] = TargetDummy
	// fmt.Println(NpcList["testNPC"])

	return NpcList
}

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

	ItemList["TestSword"] = TestSword

	return ItemList
}
