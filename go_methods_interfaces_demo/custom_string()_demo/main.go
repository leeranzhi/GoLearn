package custom_string___demo

import "fmt"

type Person struct {
	Name, Country string
}

//编写自定义String()来打印自定义字符串
func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}
func main() {
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
}
