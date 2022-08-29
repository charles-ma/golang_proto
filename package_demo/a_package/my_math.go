package some

import "fmt"

var myvar = 1
var b = c

func init() {
	fmt.Println("init my_math")
}

func Sum(i, j int) int {
	return i + j
}

func Do() {
	var v myInt = 1

	fmt.Println(v)
}

type myInt int

func (i myInt) AddOne() int {
	return int(i) + 1
}

func (i myInt) String() string {
	return "myInt" + fmt.Sprint(int(i))
}
