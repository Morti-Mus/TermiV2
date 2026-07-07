package game

type Item struct {
	Information string
	Stats       string
	Location    string
}

func Items() map[string]Item {
	return map[string]Item{
		"my-items": Item{},
	}
}
