package main

type Char struct {
	Stats
	Information
	Storage
	Location
	ObjectState
	Inventory
}

type Item struct {
	Information
	Stats
	Location
	Kind ItemKind
}

type ItemKind int

const (
	ItemKindUnknown ItemKind = iota
	ItemKindWeapon
	ItemKindPotion
)

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
	BaseDefence  int
	Defence      int
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

type Inventory struct {
	WeaponSlot [3]*Item
	PotionSlot [6]*Item
}

type ObjectState struct {
	Death map[string]Char
}
