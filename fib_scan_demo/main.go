package main

import (
	"fmt"
	"time"
)

var quitChan = make(chan bool)

func main() {
	start := time.Now()
	command := ""

	data := make(chan int, 1)
	go fib(data)
	for {
		num := <-data
		fmt.Println(num)
		_, _ = fmt.Scanf("%s", &command)
		if command == "quit" {
			quitChan <- true
			break
		}
	}
	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v second!\n", elapsed)
}

func fib(c chan<- int) {
	x, y := 1, 2
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quitChan:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}
