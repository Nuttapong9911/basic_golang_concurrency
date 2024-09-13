package main

import (
	"fmt"
	"time"
)

func main() {
	heroes := []string{"John", "Nana", "Steve", "David", "Nathan"}

	for _, hero := range heroes {
		go findTreasure(hero)
	}

	time.Sleep(2 * time.Second)
}

func findTreasure(hero string) {
	fmt.Println(hero, "start discovering the dungeon!")

	// finding the trasure
	time.Sleep(1 * time.Second)

	fmt.Println(hero, "found the treasure!")
}
