package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	heroFlags := make(chan string)
	go fightTheDragon(heroFlags)

	for {
		output, open := <-heroFlags
		if !open {
			break
		}
		fmt.Println(output)

	}

	// rounds := 5
	// for i := 0; i < rounds; i++ {
	// 	output := <-heroFlags
	// 	fmt.Println(output)
	// }

	fmt.Println("Finished")
}

func fightTheDragon(heroFlag chan string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rounds := 10
	for i := 0; i < rounds; i++ {
		// time.Sleep(time.Second)
		result := rand.Intn(10)
		heroFlag <- fmt.Sprint("Hero do the damage ", result, " hits!")
	}
	close(heroFlag)
}
