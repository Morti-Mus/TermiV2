package main

import "fmt"

func (c Char) Scan() {
	NpcList := NPC()
	// Bruh := "yo bruh"
	if c.Location.XAxis == NpcList["testNPC"].Location.XAxis && c.Location.YAxis == NpcList["testNPC"].Location.YAxis {
		fmt.Println(NpcList["testNPC"].Location.XAxis)
		fmt.Println(NpcList["testNPC"].Location.YAxis)
	}
}
