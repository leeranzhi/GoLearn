package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

//创建可重用的错误
var ErrNotFound = errors.New("err is Employee not found")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func main() {
	employee, err := getInformation(1000)
	//判断错误类型
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT Found:%v\n", err)
		return
	}
	fmt.Print(employee)
}
