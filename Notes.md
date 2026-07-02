


Idee 
1. move loop and combat loop be able to jump from one loop to the other
2. add a pick a lock minigame ? 
3. add movemnt of items inside of an array





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
3. I need to understand what and why im using it. For example in this code i have focused on maps beacuse i found the key and assignmet functionality neat for my structs. But damn its screwd me over when i need to sort or auto pick things from the maps wich you cant. Beacuse you set your own structure.

^


im running in to a problem where i want to target the map of items that is called inventory without using the key. 

I instead want to call them by possition like i will have all weapons in a certain position in the inventory array for example or a range of positions. 
So that the Iventory becomes more dynamic and i can add and lose weapons and i dont have to spesify beforhand what keys they have instead call the space in the array. 

for example weapon will always use spot 3, 4, 5 in the array. 

so im wondering if its better to transform the existing map to a slice or if i should instead re make the struct to use slice instead of map this option is something im liking a bit more. 


getting som nill pointer errors look at explenationhttps://medium.com/@ezrantn/understanding-and-preventing-nil-pointer-dereference-panics-in-go-1b0d341fa6b0

Game rules for coding 
defence will be remove /2 of t self from the damage 
crit will 2x the damage 