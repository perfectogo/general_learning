package main

import "fmt"

func main() {
	var a int

	go func() {
		a++
	}()

	//a++
	fmt.Println(a)
}
