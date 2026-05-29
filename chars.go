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
			BackPack: map[string]int{},
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
			BackPack: map[string]int{},
			Pockets:  map[string]int{},
		},
		Location: Location{
			XAxis: 0,
			YAxis: 0,
		},
	}

	NpcList["testNPC"] = TestNPC

	// fmt.Println(NpcList["testNPC"])

	return NpcList
}
