package chapter03compositetypes

import "fmt"

func Slice() {
	// var numbers []int = []int{}
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))

	//Cap(numbers)
	//Make()
	//DeclaringSlice()
	//SharingStorage()
	//ExtraConfusing()
	//Ex()
	Copy()
}

func Cap(x []int) {

	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func Make() {
	var numbers []int = make([]int, 1, 5)

	fmt.Println(numbers)
	fmt.Printf("len: %d, cap: %d\n", len(numbers), cap(numbers))

	numbers = append(numbers, 1)
	fmt.Println(numbers)
}

func DeclaringSlice() {

	var data1 []int
	fmt.Println(data1)
	fmt.Println(data1 == nil)

	var data2 = []int{}
	fmt.Println(data2)
	fmt.Println(data2 == nil)
}

func SharingStorage() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]

	x[1] = 20
	y[0] = 10
	z[1] = 30

	// When you take a slice from a slice, you are not making a copy of the data. Instead, you now have two variables that are sharing memory.
	fmt.Printf("x:, %d;\t%p\n", x, &x)
	fmt.Printf("y:, %d;\t%p\n", y, &x)
	fmt.Printf("z:, %d;\t%p\n", z, &x)
}

func ExtraConfusing() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func Ex() {

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)

	y := x[:2]
	z := x[2:]

	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	y = x[:2:2]
	z = x[2:4:4]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func Copy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

}
