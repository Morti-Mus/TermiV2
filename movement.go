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

	sum := 0
	for i := 0; i < 100; i++ {
		inp := Tester()

		if inp == "w" {
			p.Move(0, 1)
		}
		if inp == "a" {
			p.Move(-1, 0)
		}
		if inp == "d" {
			p.Move(1, 0)
		}
		if inp == "s" {
			p.Move(0, -1)
		}
		if inp == "scan" {
			p.Scan(npcList)
		}
		if inp == "scan2" {
			p.ScanV2(npcList)
		}
		if inp == "pickup" {
			// p.Pickup()
		}
		if inp == "fight" {
			fmt.Println("you have enterd fight mode")
			CombatLoop(p, npcList)
		}
		fmt.Println("\n")
		fmt.Println(p.Location.XAxis, "XAxis")
		fmt.Println(p.Location.YAxis, "YAxis")
		fmt.Printf("Your Input: ")

		sum += i
	}
}

func CombatLoop(p *Char, npcList map[string]Char) {
	a := GameDialog()
	sum := 0
	fmt.Printf(a["combat"])
	for i := 0; i < 100; i++ {
		inp := Tester()

		if inp == "attack" {
			test := FindCharsAtLocation(npcList, p.Location)

			for _, targetName := range test { // range gives index and value so keep _, other wise out put will be 0,1
				fmt.Println("Attacking:", targetName)
				p.AttackV2(npcList)
				p.npcAttack(npcList)
				fmt.Println("After attacking the enemy he strikes you back \n")
				fmt.Println(p.Stats.Health)
			}
		}
		if inp == "pickup" {
			p.PickupV2(npcList)
			fmt.Println(p.Storage.BackPack["WoddenMallet"])
		}

		if inp == "defende" {
			fmt.Println(p.Stats.BaseDefence)
			p.Defende(npcList)
			fmt.Println(p.Stats.BaseDefence)
		}

		if inp == "inspect" {
			p.inspect(npcList)
		}
		if inp == "stop" {
			return
		}
		if inp == "test" {
			p.testPickupItem(npcList)
		}
		if inp == "test2" {
			p.EquipeItem(npcList)
			fmt.Println("test")
		}
		fmt.Printf("Your Input: ")
		sum += i
	}
}

//https://medium.com/@linz07m/go-methods-88cf421c299d for learning methods

func (c *Char) Move(deltaX, deltaY int) {
	// if c.Agility < 5 {
	// 	deltaX = min(deltaX, +1)
	// 	deltaY = min(deltaY, +1)
	// }
	c.Location.XAxis += deltaX
	c.Location.YAxis += deltaY
}
