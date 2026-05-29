package main

type Char struct {
	Stats
	Information
	Storage
	Location
}

type Item struct {
}
type Information struct {
	Name        string
	Age         int
	Description string
}

type Stats struct {
	Health       int
	Mana         int
	Agility      int
	Strength     int
	Intelligence int
}

type Location struct {
	XAxis int
	YAxis int
}
type Storage struct {
	BackPack map[string]int
	Pockets  map[string]int
}
