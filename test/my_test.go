package terminusv2_test

// func (c *Char) AddItem(item *Item) error {
// 	switch item.Kind {
// 	case ItemKindWeapon:
// 		for i := range c.Inventory.WeaponSlot {
// 			if c.Inventory.WeaponSlot[i] == nil {
// 				c.Inventory.WeaponSlot[i] = item

// 				fmt.Printf(
// 					"%s was added to weapon slot %d\n",
// 					item.Name,
// 					i,
// 				)

// 				return nil
// 			}
// 		}

// 		return fmt.Errorf("all weapon slots are full")

// 	case ItemKindPotion:
// 		for i := range c.Inventory.PotionSlot {
// 			if c.Inventory.PotionSlot[i] == nil {
// 				c.Inventory.PotionSlot[i] = item

// 				fmt.Printf(
// 					"%s was added to potion slot %d\n",
// 					item.Name,
// 					i,
// 				)

// 				return nil
// 			}
// 		}

// 		return fmt.Errorf("all potion slots are full")

// 	default:
// 		return fmt.Errorf("item %q has no valid item kind", item.Name)
// 	}
// }

// func (c *Char) TestPickupItemV2(itemIndex int) {
// 	itemList := ItemArrays()

// 	if itemIndex < 0 || itemIndex >= len(itemList) {
// 		fmt.Println("Invalid item index")
// 		return
// 	}

// 	item := &itemList[itemIndex]

// 	err := c.AddItem(item)
// 	if err != nil {
// 		fmt.Println("Could not pick up item:", err)
// 	}
// }
