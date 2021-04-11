package main

import "fmt"

type Account struct {
	FullName string
	LastName string
}

func (account *Account) ChangeName(name string) {
	account.FullName = name
}

type Employee struct {
	Account
	Credits float64
}

func (e *Employee) AddCredits(credits float64) {
	e.Credits += credits
}
func (e *Employee) RemoveCredits(credits float64) {
	e.Credits -= credits
}
func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func (e Employee) String() string {
	return fmt.Sprintf("FullName is %v\nLastName is %v\nCredits:%.2f\n", e.FullName, e.LastName, e.Credits)
}

func main() {
	employee := Employee{Account{"leeranzhi", "lee"}, 23}
	fmt.Printf("before change Employee account:\n%s\n", employee.Account)
	employee.ChangeName("lee同学1988")
	fmt.Printf("Employee account:\n%s\n", employee.Account)
	fmt.Println(employee)
}
