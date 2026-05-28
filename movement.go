package main

// func Tester() string {
// 	var inpTest string
// 	_, err := fmt.Scanln(&inpTest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return inpTest
// }

// func (c Char) UseItem(itemName string, target Char) {
// 	selectedItem, has := c.BackPack[itemName]
// 	if !has {
// 		return
// 	}
// 	selectedItem.UseItemOn(target)
// }

func (c Char) Move(deltaX, deltaY int) {
	if c.Agility < 5 {
		deltaX = min(deltaX, +1)
		deltaY = min(deltaY, +1)
	}
	c.Location.XAxis += deltaX
	c.Location.YAxis += deltaY
}
