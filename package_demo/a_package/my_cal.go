package some

import "fmt"

var a = b
var c = 1

func init() {
	fmt.Println("init my_cal")
}

func Done() {
	fmt.Println(myvar, a)
	myvar = 2
}
