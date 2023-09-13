package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x, y := bits.Add(1, 2, 3)
	fmt.Println(x, y)
}
