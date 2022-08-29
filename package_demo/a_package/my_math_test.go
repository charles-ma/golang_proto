package some_test

import (
	"fmt"
	"testing"

	other "mycompany.com/package_demo/a_package"
)

func TestMyMath(t *testing.T) {
	if other.Sum(1, 2) == 3 {
		fmt.Println("good")
	}
}
