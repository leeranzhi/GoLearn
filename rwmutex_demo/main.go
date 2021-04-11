package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwLock sync.RWMutex

func read() {
	//lock.Lock()
	rwLock.Lock()
	time.Sleep(time.Millisecond)
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func write() {
	//lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
