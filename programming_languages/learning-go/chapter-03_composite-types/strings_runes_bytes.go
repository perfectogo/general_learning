package chapter03compositetypes

import "fmt"

func String() {
	var str = "😁"

	fmt.Println(str[3])
	fmt.Println(len(str))

	for _, i := range str {
		fmt.Println("\n---",string(i))
	}

}