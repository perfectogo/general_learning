package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var ch = make(chan int)

	go GetRandomNumber(ch)

	for rNum := range ch {
		fmt.Println("random number:", rNum)
	}
}

func GetRandomNumber(ch chan int) {
	rand.Seed(time.Now().UnixNano())

	for {
		num := rand.Intn(100)
		time.Sleep(1 * time.Second)

		ch <- num
	}
	//close(ch)
}
