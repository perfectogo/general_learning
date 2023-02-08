package main

import "fmt"

type Car struct {
	name string
}

func MakeCar(name string) Car {

	if name != "" {
		return Car{name}
	}
	return Car{}
}

func main() {
	car := MakeCar("Malibu")

	fmt.Println(car)
}
