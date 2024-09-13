package main

import (
	"fmt"
	"sync"
	"time"
)

var ready bool

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	workerNums := 3
	wg.Add(workerNums)

	go works(cond, &wg, "Smith")
	go works(cond, &wg, "Steve")
	go works(cond, &wg, "Salami")

	// precessing and change the condtion
	fmt.Println("Data is processing.")
	time.Sleep(5 * time.Second)
	ready = true

	// broadcasting
	fmt.Println("Data is ready, start broadcasting")
	cond.Broadcast()

	fmt.Println("Wait for all workers")
	wg.Wait()

}

func works(cond *sync.Cond, wg *sync.WaitGroup, name string) {
	time.Sleep(time.Second)

	cond.L.Lock()
	for {
		if ready {
			fmt.Println("Worker", name, " has done the job.")
			wg.Done()
			break
		}
		fmt.Println("Worker ", name, " is waiting.")
		cond.Wait()
	}
	cond.L.Unlock()
}
