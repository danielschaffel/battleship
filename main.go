package main

import (
	"fmt"
	"ship/game"
)

func main() {
	p := game.NewPlayer()
	var added, err = p.AddBoat(0, 0, 2, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added a boat %t\n", added)
	fmt.Printf("%v\n", p)

	added, err = p.AddBoat(0, 0, 2, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added a boat %t\n", added)
	fmt.Printf("%v\n", p)
}
