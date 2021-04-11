package main

import "fmt"

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		if num, ok := <-ch1; ok {
			ch2 <- num * num
			continue
		}
		break
	}
	//for num := range ch1 {
	//	ch2 <- num * num
	//}
	close(ch2)
}

func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
