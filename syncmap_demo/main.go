package main

import (
	"fmt"
	"sync"
)

type Student struct {
	name string
	age  int
}

//sync.map 并发安全的map
var m = make(map[int]int)
var wg sync.WaitGroup
var m2 sync.Map

var student Student = Student{}

func get(key int) int {
	return m[key]
}
func set(key, value int) {
	m[key] = value
}
func main() {
	fmt.Println(student)
	//for i := 0; i < 20; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		//set(i, i+100)
	//		m2.Store(i, i+100)
	//		value, _ := m2.Load(i)
	//		fmt.Printf("key:%v,valye:%v\n", i, value)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
}
