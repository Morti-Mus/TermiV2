package main

import (
	"fmt"
	"math/rand/v2"
)

func (c Char) Scan(NpcList map[string]Char) {
	// dialog := GameDialog()

	for i, Npc := range NpcList {

		fmt.Printf("%v x%v/y%v \n",
			i,
			Npc.Location.XAxis,
			Npc.Location.YAxis,
		)
	}
}

func (c *Char) EquipeItem(NpcList map[string]Char) { // https://leapcell.io/blog/checking-if-a-key-exists-in-a-go-map
	test := c.Storage.BackPack

	for key, value := range test {
		fmt.Println("WeaponName: ", key)
		c.Stats.Strength += value.Stats.Strength
		break
	}
}

func (c *Char) PickupV2(NpcList map[string]Char) {
	targets := FindCharsAtLocation(NpcList, c.Location)

	if len(targets) == 0 {
		fmt.Println("No availabe targets try scan")
	}

	targetName := targets[0]
	loot := make(map[string]Item)

	loot["testHammer"] = NpcList[targetName].Storage.LootTable["testHammer"]

	if NpcList[targetName].Stats.Health <= 0 {
		c.Storage.BackPack["testHammer"] = loot["testHammer"]
	}
}

func (c *Char) TestPickupItem(NpcList map[string]Char) {
	itemList := Items()

	c.Storage.BackPack["testHammer"] = itemList["testHammer"]
	fmt.Println(c.Storage.BackPack["testHammer"])
}

// func (c Char) PickUpP() {
// 	ItemList := ItemArrays()

//		fmt.Println(ItemList[])
//	}

func (c *Char) UsePotion(index int) {
	for i := range c.Inventory.PotionSlot {
		if c.Inventory.PotionSlot[i] != nil {
			fmt.Println(i, c.Inventory.PotionSlot[i].Name)
		}
	}
	// need to add max health cap
	Potion := c.Inventory.PotionSlot[index]

	c.Health += Potion.Health
	c.Mana += Potion.Mana
	c.Agility += Potion.Agility
	c.Strength += Potion.Strength
	c.Intelligence += Potion.Intelligence

}

func (c *Char) TestPickupItemV2(index int) {
	ItemList := ItemArrays()

	// fmt.Print(ItemList[0])

	// if c.Inventory.WeaponSlot[0] != nil {
	// 	ItemList = append(ItemList, *c.Inventory.WeaponSlot[0])
	// }
	// c.Inventory.WeaponSlot[0] = &weapons
	// c.Inventory.WeaponSlot[1] = &weapons

	c.Inventory.WeaponSlot[index] = &ItemList[index]
	c.Inventory.PotionSlot[index] = &ItemList[index]
	// c.Inventory.WeaponSlot[index] = &ItemList[index]
	// c.Inventory.WeaponSlot[2] = &ItemList[index]

	fmt.Println("test", c.Inventory.WeaponSlot[0])
	fmt.Println("test", c.Inventory.WeaponSlot[1])
	fmt.Println("test", c.Inventory.WeaponSlot[2])
	fmt.Println("Potion", c.Inventory.PotionSlot[0])
}
func (c *Char) TestPickupItemV3(itemIndex int) {
	itemList := ItemArrays()

	if itemIndex < 0 || itemIndex >= len(itemList) {
		fmt.Println("Invalid item index")
		return
	}

	item := &itemList[itemIndex]

	err := c.AddItem(item)
	if err != nil {
		fmt.Println("Could not pick up item:", err)
	}
}
func (c *Char) AddItem(item *Item) error {
	switch item.Kind {
	case ItemKindWeapon:
		for i := range c.Inventory.WeaponSlot {
			if c.Inventory.WeaponSlot[i] == nil {
				c.Inventory.WeaponSlot[i] = item

				fmt.Printf(
					"%s was added to weapon slot %d\n",
					item.Name,
					i,
				)

				return nil
			}
		}

		return fmt.Errorf("all weapon slots are full")

	case ItemKindPotion:
		for i := range c.Inventory.PotionSlot {
			if c.Inventory.PotionSlot[i] == nil {
				c.Inventory.PotionSlot[i] = item

				fmt.Printf(
					"%s was added to potion slot %d\n",
					item.Name,
					i,
				)

				return nil
			}
		}

		return fmt.Errorf("all potion slots are full")

	default:
		return fmt.Errorf("item %q has no valid item kind", item.Name)
	}
}

func (c *Char) DisplayInventory() {

	testTwo := c.Inventory

	for i := 0; i < len(testTwo.PotionSlot); i++ {
		if testTwo.PotionSlot[i] != nil {
			fmt.Println(testTwo.PotionSlot[i].Name)
			// continue
		}
		// fmt.Println(testTwo[i].Name)
	}

	for i := 0; i < len(testTwo.WeaponSlot); i++ {
		if testTwo.WeaponSlot[i] != nil {
			fmt.Println(testTwo.WeaponSlot[i].Name)
			// continue
		}
		// fmt.Println(testTwo[i].Name)
	}
}

func (c *Char) Defende(NpcList map[string]Char) {

	// c.Stats.BaseDefence = c.Stats.Strength + c.Stats.BaseDefence
	c.Stats.Defence = c.BaseDefence + c.Stats.Strength

}

func (c Char) Inspect(NpcList map[string]Char) {
	targets := FindCharsAtLocation(NpcList, c.Location)

	if len(targets) == 0 {
		fmt.Println("No available targets try scan")
		return
	}
	targetName := targets[0]
	fmt.Println("you have encounterd: ", NpcList[targetName].Name, "\n")
	fmt.Println("after a closer inspection you se: ", NpcList[targetName].Description)
	fmt.Println("----------------------")

	fmt.Println("his health is", NpcList[targetName].Health)
	fmt.Println("his mana is ", NpcList[targetName].Stats.Mana)
	fmt.Println("his strength is", NpcList[targetName].Stats.Strength)
	fmt.Println(NpcList[targetName].Stats.Mana)
}

func (c Char) AttackV2(NpcList map[string]Char) Char { // need to redo attack so it dosent create new copy and works on the closest npc

	targets := FindCharsAtLocation(NpcList, c.Location)

	if len(targets) == 0 {
		fmt.Println("No availabe targets try scan")
		return c
	}

	targetName := targets[0]

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

func (c Char) NpcDefence(NpcList map[string]Char) int {
	attacker := FindCharsAtLocation(NpcList, c.Location)

	if len(attacker) == 0 {
		fmt.Println("No available attacker")
	}

	attackerName := attacker[0]

	test := NpcList[attackerName]
	test.Stats.Defence = test.BaseDefence + test.Stats.Strength
	NpcList[attackerName] = test
	return NpcList[attackerName].Stats.Defence
}

func (c *Char) NpcAttack(NpcList map[string]Char) int {
	attacker := FindCharsAtLocation(NpcList, c.Location)

	if len(attacker) == 0 {
		fmt.Println("No available attacker")
	}

	attackerName := attacker[0]

	var baseDmgCalc int = NpcList[attackerName].BaseDmg + NpcList[attackerName].Strength
	baseDmgCalc -= c.Stats.Defence / 2
	c.Stats.Health = c.Stats.Health - baseDmgCalc
	return baseDmgCalc
}

func (c Char) NpcDeath(NpcList map[string]Char) {
	attacker := FindCharsAtLocation(NpcList, c.Location)

	if len(attacker) == 0 {
		fmt.Println("No available attacker")
		return
	}

	attackerName := attacker[0]

	NpcList[attackerName].ObjectState.Death[""] = NpcList[attackerName]

	delete(NpcList, attackerName)
}

func (c *Char) NpcChoiceAction(NpcList map[string]Char) {
	Defender := FindCharsAtLocation(NpcList, c.Location)

	if len(Defender) == 0 {
		fmt.Println("No available attacker")
		return
	}

	DefenderName := Defender[0]

	RandInt := rand.IntN(100)

	if RandInt < 65 {
		c.NpcAttack(NpcList)
		fmt.Println("After attacking the enemy he strikes you back")
		fmt.Println("You take: ", c.NpcAttack(NpcList), "Dmg")
		fmt.Println("Remaining Health", c.Stats.Health)
	} else if RandInt < 100 {
		fmt.Println("Then enemy is boulstering their defence")
		c.NpcDefence(NpcList)
		fmt.Println("The Npc defence rose by", NpcList[DefenderName].Stats.Defence)
	}

}

func (c Char) ItemDrop(NpcList map[string]Char) map[string]Item {
	itemList := Items()
	npcTargetList := FindCharsAtLocation(NpcList, c.Location)
	npcCurrentTarget := npcTargetList[0]

	loot := NpcList[npcCurrentTarget].Storage.LootTable
	loot["WoodenMallet"] = itemList["WoodenMallet"]

	return loot
}

func SamePosition(a Location, b Location) bool {
	return a.XAxis == b.XAxis && a.YAxis == b.YAxis
}

func FindCharsAtLocation(chars map[string]Char, target Location) []string { // Check if  we can use pointers instead of copy logic
	matches := []string{}

	for name, char := range chars {
		if SamePosition(char.Location, target) {
			matches = append(matches, name)
		}
	}
	return matches
}

func (c *Char) Move(deltaX, deltaY int) {
	// if c.Agility < 5 {
	// 	deltaX = min(deltaX, +1)
	// 	deltaY = min(deltaY, +1)
	// }
	c.Location.XAxis += deltaX
	c.Location.YAxis += deltaY
}

func GameDialog() map[string]string {
	a := make(map[string]string)

	a["Intro"] = "Hello welcome to the game you will soon get instructions on how it works \n"
	a["IntroWeaponPick"] = "Pick upp your weapon of choice \n (1) for starter Sword \n (2) for starter Axe \n (3) for starter Mace\n"
	a["IntroPotionsPick"] = "Pick upp your potion of choice \n (4) for Health potion \n (5) for Mana potion \n"
	a["Movment"] = "You will be able to move in the cardinal directions. \n (w) for north \n (a) for east \n (d) for west \n (s) for south \n"
	a["Attack"] = "stuff"
	a["Scan"] = "you see a person in the distance \n"
	a["Combat"] = "These are the options you have in combat. \n (attack) to attack your enemy \n (inspect) to inspect you enemy \n (defende) to defende from your enemy \n (stop) to return to movment.\n"

	return a
}
