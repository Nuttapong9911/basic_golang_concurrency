package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var army sync.WaitGroup

	start := time.Now()

	army.Add(3)
	go prepareTheHero("Alex", &army)
	go prepareTheHero("Doe", &army)
	go prepareTheHero("Nick", &army)
	// go lazyPrepareTheHero("Syphon", &army)

	time.Sleep(5 * time.Second)
	fmt.Println("Waiting")
	army.Wait() // add == 0 ?

	fmt.Println(time.Since(start))

	fmt.Println("Heroes are ready for the mission.")
}

func prepareTheHero(name string, army *sync.WaitGroup) {
	fmt.Println("Start traing the hero", name)

	// train the hero
	time.Sleep(2 * time.Second)
	fmt.Println("Hero", name, "is now strong!")

	// finished training
	army.Done() // add -= 1
}

func lazyPrepareTheHero(name string, army *sync.WaitGroup) {
	fmt.Println("Start traing the hero", name)
	time.Sleep(5 * time.Second)
	fmt.Println("Hero", name, "is now strong?")

	// what if he quit the trainging?
	army.Done()
}
