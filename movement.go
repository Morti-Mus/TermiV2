package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Tester() string {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	return strings.TrimSpace(inp)
}

// https://medium.com/@danielabatibabatunde1/pointers-in-golang-240b30c6940d for learning pointers
func MoveLoop(p *Char, NpcList map[string]Char) {

	a := GameDialog()

	fmt.Printf(a["Intro"])

	fmt.Printf(a["Movment"])
	fmt.Printf("Your Input: ")
	// p := &mainChar // p is the pointer for mainChars address in memory.

	for {
		inp := Tester()

		switch inp {

		case "w":
			p.Move(0, 1)

		case "a":
			p.Move(-1, 0)

		case "d":
			p.Move(1, 0)

		case "s":
			p.Move(0, -1)

		case "scan":
			p.Scan(NpcList)

		case "test":
			p.TestPickupItemV2(0)

		case "inventory":
			p.InventoryLoop()

		case "fight":
			fmt.Println("you have enterd fight mode")
			CombatLoop(p, NpcList)
		}

		fmt.Println("\n")
		fmt.Println(p.Location.XAxis, "XAxis")
		fmt.Println(p.Location.YAxis, "YAxis")
		fmt.Printf("Your Input: ")
	}
}

func (c *Char) InventoryLoop() {
	a := GameDialog()

	fmt.Printf(a["IntroWeaponPick"])
	fmt.Printf(a["IntroPotionsPick"])
	for {
		inp := Tester()

		fmt.Println()
		switch inp {

		case "1": // add user fmtprint to explain weapon choice
			c.TestPickupItemV3(0)

		case "2":
			c.TestPickupItemV3(1)

		case "3":
			c.TestPickupItemV3(2)

		case "4":
			c.TestPickupItemV3(3)

		case "5":
			c.TestPickupItemV3(4)
		case "6": //Create subb menu for potion usage and item equipment will get complicated
			c.EquipeItemV2(0)

		case "inv":
			c.DisplayInventory()

		case "try":
			fmt.Println(c.Stats.Health)
			c.UsePotion(0)
			fmt.Println(c.Stats.Health)
		case "test":
			fmt.Println("test")

		case "stop":
			return
		}

	}
}

func CombatLoop(c *Char, NpcList map[string]Char) {
	a := GameDialog()
	c.Stats.Defence = c.Stats.BaseDefence
	// fmt.Println(c.Stats.Defence)
	fmt.Printf(a["Combat"])

	for {
		inp := Tester()

		switch inp {

		case "attack":
			test := FindCharsAtLocation(NpcList, c.Location)

			for _, targetName := range test { // range gives index and value so keep _, other wise out put will be 0,1
				fmt.Println("Attacking:", targetName, "\n")
				c.AttackV2(NpcList)
				c.NpcChoiceAction(NpcList)
				// c.NpcAttack(NpcList)
				// fmt.Println("After attacking the enemy he strikes you back")
				// fmt.Println("You take: ", c.NpcAttack(NpcList), "Dmg")
				// fmt.Println("Remaining Health", c.Stats.Health)
			}

		case "pickup":
			c.PickupV2(NpcList)
			fmt.Println(c.Storage.BackPack["WoddenMallet"])

		case "defende":
			// fmt.Println(p.Stats.BaseDefence)
			c.Defende(NpcList)
			c.NpcChoiceAction(NpcList)
			fmt.Println(c.Stats.Defence)

		case "inspect":
			c.Inspect(NpcList) // check if cases are alla evaluated or not

		case "stop":
			return

		case "test":
			c.TestPickupItem(NpcList)

		case "test2":
			fmt.Println(c.Stats.Strength)
			c.EquipeItem(NpcList)
			fmt.Println(c.Stats.Strength)

		}
		fmt.Printf("Your Input: ")
	}
}

//https://medium.com/@linz07m/go-methods-88cf421c299d for learning methods
