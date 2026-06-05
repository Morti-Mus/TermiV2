package main

import (
	"fmt"
	"math/rand/v2"
)

func (c Char) Scan() {
	NpcList := NPC()
	dialog := GameDialog()

	fmt.Println(NpcList["testNPC"].Location.XAxis, dialog["scan"])
	fmt.Println(NpcList["testNPC"].Location.YAxis, dialog["scan"])
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

func (c Char) AttackV2(NpcList map[string]Char) Char { // need to redo attack so it dosent create new copy and works on the closest npc
	targets := FindCharsAtLocation(NpcList, c.Location)

	if len(targets) == 0 {
		fmt.Println("No availabe targets try scan")
	}

	targetName := targets[0]
	fmt.Println("you have encounterd: ", NpcList[targetName].Name, "\n")
	fmt.Println("after a closer inspection you se: ", NpcList[targetName].Description)
	fmt.Println("----------------------")
	fmt.Println("his health is", NpcList[targetName].Health)

	var baseDmgCalc int = c.Stats.Strength + c.BaseDmg
	if CritChanse := rand.IntN(100 - c.Stats.Agility); CritChanse == 0 {
		damage := baseDmgCalc * 2

		npc := NpcList[targetName]
		npc.Stats.Health -= damage
		NpcList[targetName] = npc

		fmt.Println(damage)
		fmt.Println(NpcList[targetName].Stats.Health)

		return npc
	}

	npc := NpcList[targetName]
	npc.Stats.Health -= baseDmgCalc
	NpcList[targetName] = npc

	fmt.Println("you hit him for", baseDmgCalc)
	fmt.Println(NpcList[targetName].Stats.Health, "enemy health remaining")

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

// func (c Char) Defende() {
// 	var baseDefenceCalc int = c.Stats.Strength + c.BaseDefence

// }

func GameDialog() map[string]string {
	a := make(map[string]string)

	a["intro"] = "Hello welcome to the game you will soon get instructions on how it works \n"
	a["movment"] = "You will be able to move in the cardinal directions. \n (w) for north \n (a) for east \n (d) for west \n (s) for south"
	a["attack"] = "stuff"
	a["scan"] = "you see a person in the distance \n"
	return a
}
