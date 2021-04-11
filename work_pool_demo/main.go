package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//开启三个goroutine
	for i := 0; i < 3; i++ {
		go work(i, jobs, result)
	}

	//发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//deadlock
	//for ret := range result {
	//	fmt.Println(ret)
	//}
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
}

func work(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		resultNum := job * 2
		time.Sleep(5 * time.Microsecond)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
		result <- resultNum
	}
}
