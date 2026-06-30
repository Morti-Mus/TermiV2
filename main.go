package main

func main() {
	// test()
	// NPC()
	GameLoops()
}

func GameLoops() {
	// MoveLoop()

	mainChar := PlayerChar()
	p := &mainChar
	npcList := NPC()

	MoveLoop(p, npcList) // <--- iwant to move them here to have them in a higher scope so they dont get re called.
}
