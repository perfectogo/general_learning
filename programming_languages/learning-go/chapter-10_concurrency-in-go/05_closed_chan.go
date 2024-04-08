package main

import "fmt"

func Run(ch chan int) {
	ch <- 7
	close(ch)
}
func main() {

	ch := make(chan int)

	go Run(ch)

	d, ok := <-ch
	fmt.Println(d, ok)

	d, ok = <-ch
	fmt.Println(d, ok)

}
