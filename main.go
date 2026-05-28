package main

import "fmt"

func main() {

}

func Tester() string {
	var inpTest string
	_, err := fmt.Scanln(&inpTest)
	if err != nil {
		fmt.Println(err)
	}

	return inpTest
}

func GameLoop() {
	inp := Tester()

}
