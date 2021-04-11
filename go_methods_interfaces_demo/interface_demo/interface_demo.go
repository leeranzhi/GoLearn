package interface_demo

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}
type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * s.size
}

func (s Square) Area() float64 {
	return s.size * s.size
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.radius
}

/**
优点：对于Shape的每个新类型或者实现，printInfo函数都不需要更改，代码变得灵活，易扩展
*/
func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func main() {
	var s Shape = Square{3}
	printInformation(s)
	c := Circle{6}
	printInformation(c)
}
