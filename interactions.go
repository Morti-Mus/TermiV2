package main

import (
	"fmt"
	"math/rand/v2"
)

func (c Char) Scan() {
	NpcList := NPC()
	if c.Location.XAxis == NpcList["testNPC"].Location.XAxis && c.Location.YAxis == NpcList["testNPC"].Location.YAxis {
		fmt.Println(NpcList["testNPC"].Location.XAxis)
		fmt.Println(NpcList["testNPC"].Location.YAxis)
	}
}

func (c Char) Pickup() {
	ItemList := Items()

	if c.Location.XAxis == ItemList["TestSword"].Location.XAxis && c.Location.YAxis == ItemList["TestSword"].Location.YAxis {
		c.Storage.BackPack["TestSword"] = ItemList["TestSword"]
		fmt.Println(ItemList, "List befor deletion")
		delete(ItemList, "TestSword")
		fmt.Println(ItemList, "List after deletion")
	}
	fmt.Println(c.Storage.BackPack["TestSword"])
}

// func (c Char) EnterFightModeLoop() {
// 	NpcList := NPC()
// 	if c.Location.XAxis == NpcList["testNPC"].Location.XAxis && c.Location.YAxis == NpcList["testNPC"].Location.YAxis {
// 		CombatLoop()
// 	}
// }

func (c Char) Attack() Char {
	NpcList := NPC() //  to auto pick the npc that you are closest to dont know how tho
	var baseDmgCalc int = c.Stats.Strength + c.BaseDmg
	if CritChanse := rand.IntN(100 - c.Stats.Agility); CritChanse == 0 {
		damage := baseDmgCalc * 2

		npc := NpcList["testNPC"]
		npc.Stats.Health -= damage
		NpcList["testNPC"] = npc

		return npc

	}
	npc := NpcList["testNPC"]
	npc.Stats.Health -= baseDmgCalc
	NpcList["testNPC"] = npc
	return npc
}

func SamePosition(a Location, b Location) bool {
	return a.XAxis == b.XAxis && a.YAxis == b.YAxis
}

func FindCharsAtLocation(chars map[string]Char, target Location) []string { // need to look att this more still feels like magic
	matches := []string{}

	for name, char := range chars {
		if SamePosition(char.Location, target) {
			matches = append(matches, name)
		}
	}
	return matches
}
