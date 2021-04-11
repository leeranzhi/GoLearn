package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	size := 15
	resultChan := make(chan float64, size)

	for i := 0; i < size; i++ {
		go fib(float64(i), resultChan)
	}
	for i := 0; i < size; i++ {
		fmt.Printf("Fib(%v): %v\n", i, <-resultChan)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v second!\n", elapsed)
}

func fib(number float64, resultChan chan<- float64) {
	x, y := 1.0, 2.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	resultChan <- x
}
