package main

type Char struct {
	Stats
	Information
	Storage
	Location
	ObjectState
}

type Item struct {
	Information
	Stats
	Location
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
	BaseDmg      int
}

type Location struct {
	XAxis int
	YAxis int
}
type Storage struct {
	BackPack  map[string]Item
	Pockets   map[string]int
	LootTable map[string]Item
}

type ObjectState struct {
	Death map[int]int
}
