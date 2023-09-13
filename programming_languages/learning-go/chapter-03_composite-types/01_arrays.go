package chapter03compositetypes

import (
	"fmt"
)

func Arrays() {
	var numbers = [...]int{}
	//numbers= append(numbers, 1) // we cannot appent new element for arrays numbers
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	var numbers2 [5]int = [5]int{1, 2, 3}
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
}


