package game_test

import (
	"testing"

	"github.com/Morti-Mus/TermiV2.git/game"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		itemName string
	}{
		{
			desc:     "hello",
			itemName: "my-item",
		},
		{
			desc:     "hello",
			itemName: "my-items",
		},
		{
			desc:     "hello",
			itemName: "your-items",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			items := game.Items()
			_, ok := items[tC.itemName]
			t.Log("testing", tC.itemName, tC.desc)
			if ok != false {
				t.Errorf("the items %s exsist as value", tC.itemName)
			}
		})
	}
}
