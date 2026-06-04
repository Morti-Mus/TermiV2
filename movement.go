package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func Tester() string {
// 	var inpTest string
// 	_, err := fmt.Scanln(&inpTest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return inpTest
// }

//	func (c Char) UseItem(itemName string, target Char) {
//		selectedItem, has := c.BackPack[itemName]
//		if !has {
//			return
//		}
//		selectedItem.UseItemOn(target)
//	}
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
			p.Scan() // will break this out to its own loop but have to figure things out
		}
		if inp == "pickup" {
			p.Pickup()
		}
		if inp == "fight" {
			fmt.Println("you have enterd fight mode")
			CombatLoop(p, npcList)
		}
		fmt.Println("\n")
		fmt.Println(p.Location.XAxis)
		fmt.Println(p.Location.YAxis)
		// fmt.Println(*p) // i think i should print the p if i want to se the changes
		// fmt.Println(p)
		sum += i
	}
}

func CombatLoop(p *Char, npcList map[string]Char) {
	// MainChar := PlayerChar()
	NpcList := NPC()
	// n := &NpcList
	// p := &MainChar

	sum := 0
	for i := 0; i < 100; i++ {
		inp := Tester()

		if inp == "Attack" && p.Location.XAxis == NpcList["testNPC"].Location.XAxis && p.Location.YAxis == NpcList["testNPC"].Location.YAxis {
			p.Attack()
		}
		if inp == "test" {
			test := FindCharsAtLocation(NpcList, p.Location)

			for _, targetName := range test {
				fmt.Println("Attacking:", targetName)
				p.Attack()
			}
		}

		if inp == "stop" {
			return
		}

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
