package main

import (
	"fmt"
	"time"
)

func main() {

	var chan1 = make(chan string)
	var chan2 = make(chan string)

	go func() {
		for {
			chan1 <- "speedy"
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			chan2 <- "slowly"
			//time.Sleep(1 * time.Second)
		}
	}()

	// for {
	// 	fmt.Println(<-chan1)
	// 	fmt.Println(<-chan2)
	// }

	for {
		select {
		case m1 := <-chan1:
			fmt.Println(m1)
		case m2 := <-chan2:
			fmt.Println(m2)
		default:
			fmt.Println("ma'lumot yo'q")

		}

	}
}
