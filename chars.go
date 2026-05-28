package main

func PlayerChar(Char) {
	MainPlayer := Char{
		information: information{
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
	}
}
