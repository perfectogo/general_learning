package main

import "fmt"

type calculator struct {
}

type calculatorOperationsI interface {
	addition(x, y float64) float64
	subtraction(x, y float64) float64
	multiplication(x, y float64) float64
	division(x, y float64) float64
}

func (c calculator) addition(x, y float64) float64 {
	return x + y
}

func (c calculator) subtraction(x, y float64) float64 {
	return x - y

}

func (c calculator) multiplication(x, y float64) float64 {
	return x * y
}

func (c calculator) division(x, y float64) float64 {
	return x / y
}

func Calculate(x, y int) calculatorOperationsI {
	result := calculator{}
	return result
}

func main() {
	calc := calculator{}

	fmt.Println(calc.addition(5, 7))
}
