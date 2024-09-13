package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	heroFlag1, heroFlag2 := make(chan interface{}), make(chan interface{})

	go doTheMission(heroFlag1, "Edward")
	go doTheMission(heroFlag2, "Arthur")

	select {
	case message := <-heroFlag1:
		fmt.Println(message)
		fmt.Println("Get back to the kingdom!")
	case message := <-heroFlag2:
		fmt.Println(message)
		fmt.Println("Get back to the kingdom!")
	}

	fmt.Println("Finished")
}

func doTheMission(channel chan interface{}, hero string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// doing the mission
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	channel <- fmt.Sprint("Hero ", hero, " failed the mission!")
}
