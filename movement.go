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
func MoveLoop(p *Char, npcList map[string]Char) {

	a := GameDialog()

	fmt.Printf(a["intro"])

	fmt.Printf(a["movment"])
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
			p.Scan(npcList)

		case "scan2":
			p.ScanV2(npcList)

		case "pickup":
			// p.Pickup()

		case "fight":
			fmt.Println("you have enterd fight mode")
			CombatLoop(p, npcList)
		}

		fmt.Println("\n")
		fmt.Println(p.Location.XAxis, "XAxis")
		fmt.Println(p.Location.YAxis, "YAxis")
		fmt.Printf("Your Input: ")
	}
}

func CombatLoop(p *Char, npcList map[string]Char) {
	a := GameDialog()

	fmt.Printf(a["combat"])

	for {
		inp := Tester()

		switch inp {

		case "attack":
			test := FindCharsAtLocation(npcList, p.Location)

			for _, targetName := range test { // range gives index and value so keep _, other wise out put will be 0,1
				fmt.Println("Attacking:", targetName, "\n")
				p.AttackV2(npcList)
				p.npcAttack(npcList)
				fmt.Println("After attacking the enemy he strikes you back")
				fmt.Println("You take: ", p.npcAttack(npcList), "Dmg")
				fmt.Println("Remaining Health", p.Stats.Health)
			}

		case "pickup":
			p.PickupV2(npcList)
			fmt.Println(p.Storage.BackPack["WoddenMallet"])

		case "defende":
			fmt.Println(p.Stats.BaseDefence)
			p.Defende(npcList)
			fmt.Println(p.Stats.BaseDefence)

		case "inspect":
			p.inspect(npcList)

		case "stop":
			return

		case "test":
			p.testPickupItem(npcList)

		case "test2":
			p.EquipeItem(npcList)
			fmt.Println("test")

		}
		fmt.Printf("Your Input: ")
	}
}

//https://medium.com/@linz07m/go-methods-88cf421c299d for learning methods
