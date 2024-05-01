package main

import "fmt"

func main() {

	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearch(numbers, 2))
}

func binarySearch(list []int, item int) any {
	var (
		low  = 0
		high = len(list) - 1
	)

	for low <= high {
		mid := low + high

		gues := list[mid]

		if gues == item {
			return mid
		}

		if gues > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nil
}
