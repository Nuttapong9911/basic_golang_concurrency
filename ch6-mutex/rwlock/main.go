package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	rwLock sync.RWMutex
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	go read()
	go read()
	go read()
	go read()
	go write()
	go read()
	go read()
	go write()

	time.Sleep(6 * time.Second)
}

func read() {
	random := rand.Intn(5)
	time.Sleep(time.Duration(random) * time.Second)

	rwLock.RLock()
	fmt.Println("Read is Locked")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Read is Unlocked")
	rwLock.RUnlock()
}

func write() {
	random := rand.Intn(5)
	time.Sleep(time.Duration(random+1) * time.Millisecond)

	rwLock.Lock()
	fmt.Println("Write is Locked")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Write is Unlocked")
	rwLock.Unlock()
}
