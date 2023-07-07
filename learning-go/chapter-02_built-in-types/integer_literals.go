package chapter02builtintypes

import "fmt"

func Integer() {
	var number int = 1_000_000_000
	fmt.Println(number)
}

func IntegerType() {

}

func SpecialIntegerTypes() {
	var (
		first  byte = 232 // uint8
		second int  = 999999999999999999
		third  uint = 9999999999999999999
	)

	fmt.Println(first, second, third)
}

func IntegerOperators() {

	a := 3 << 4 << 3 << 3

	fmt.Println(a)
}
