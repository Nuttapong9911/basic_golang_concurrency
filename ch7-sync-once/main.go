package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	rand.New(rand.NewSource(time.Now().UnixNano()))

	iterations := 1000
	sum := 0
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			// sum++
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			once.Do(func() {
				fmt.Println("Done by routine#", i)
				sum++
			})

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Finished sum:", sum)
}
