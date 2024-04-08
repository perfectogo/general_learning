package main

import (
	"fmt"
	"time"
)

func main() {

	go run("hello, world")
	go run("bye")
}

func run(s string) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
