


Idee 
1. move loop and combat loop be able to jump from one loop to the other





Game Plan. 
1. start game have intro on what different commands you can write. 
2. pick upp items in that intro.
3. go throug combat tutorial vs target dummy. 



Things i want to add.
1. add inspect to se enemy stats. 
2. first fight vs target dummy. 
3. fight logic where target dummy can pick between attack defend or retreat where different choicses are randomised but have higher lower chance. 
4. need to figure out how to transfer dmg and stuff when combat mode is trigger





		damage := baseDmgCalc * 2
		if npc, ok := NpcList["testNPC"]; ok {
			npc.Stats.Health -= damage
			NpcList["testNPC"] = npc
		}


Lessons learned! 
1. i need to move the creation of all variables that are gonna be persistant to main()
Then use pointers to manipulet them in the loops so they will never be re called in the code. 

2. Call order goes main() whats inside main and so on. So from the top down and not bottom upp. 
this is a bitt of a important concept i think 