package methods

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

type colorTriangle struct {
	triangle
	color string
}

//重载colorTriangle 中 triangle 的 perimeter方法
func (t colorTriangle) perimeter() int {
	return t.size * 6
}

//声明其他类型的方法
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	t := colorTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter(colored) ", t.perimeter())         //此处调用colorTriangle 中 triangle的方法，我们也可以重载此方法
	fmt.Println("Perimeter(normal) ", t.triangle.perimeter()) //也可以访问"原始"方法

}
