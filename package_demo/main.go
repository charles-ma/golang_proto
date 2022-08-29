package main

import (
	"fmt"

	some "mycompany.com/package_demo/a_package"
)

func main() {
	fmt.Println(some.Sum(1, 2))

	some.Do()
	some.Done()
}
