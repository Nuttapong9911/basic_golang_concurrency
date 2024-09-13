package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count int
)

func main() {
	heroes := 1000
	for i := 0; i < heroes; i++ {
		go storeTheTreasure()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Now we have count:", count)
}

func storeTheTreasure() {
	lock.Lock()
	count++ // count = count + 1
	// temp := count
	// temp = temp+ 1
	// count = temp
	lock.Unlock()
}
