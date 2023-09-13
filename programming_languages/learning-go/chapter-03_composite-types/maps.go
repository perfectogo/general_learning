package chapter03compositetypes

import "fmt"

func CommaOk() {
	var word = make(map[string]string, 1)

	word["go"] = "went"

	if value, ok := word["go"]; ok {
		fmt.Println(value)
	}

	if value, ok := word["went"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("none")
	}

	fmt.Println(word)

	Delete(word)

	fmt.Println(word)
}

func Delete(word map[string]string) {
	delete(word, "go")
}
