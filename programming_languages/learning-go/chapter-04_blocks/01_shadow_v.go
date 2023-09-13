package chapter04blocks

import (
	"fmt"
)

func Shadow() {
	var x int = 10

	func() {
		fmt.Println(x)

		var x int = 5 // shadowing variable

		fmt.Println(x)
	}()

	fmt.Println(x)
}
