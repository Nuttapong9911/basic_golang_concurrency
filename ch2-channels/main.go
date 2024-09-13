package main

import (
	"fmt"
	"time"
)

func main() {
	signalFlag := make(chan bool, 1)
	hero := "Juliot"

	go fightTheDragon(hero, signalFlag)

	// go spyFlag(signalFlag)

	time.Sleep(3 * time.Second)
	fmt.Println("The king wait for the signal")
	fmt.Println("The king got the flag", <-signalFlag)
	fmt.Println("What a peaceful day!")
}

// child routine
func fightTheDragon(hero string, signalFlag chan bool) {
	fmt.Println("Hero", hero, "face the Dragon!")

	// do his best to win the dragon
	time.Sleep(1 * time.Second)
	fmt.Println("Hero", hero, "win against the Dragon, Hurrey!")

	// tell the king that he won
	signalFlag <- true
	fmt.Println("Hero", hero, "already sent the flag")

	// time.Sleep(10 * time.Second)
}

func spyFlag(signalFlag chan bool) {
	fmt.Println("Spy got the flag", <-signalFlag)
}
