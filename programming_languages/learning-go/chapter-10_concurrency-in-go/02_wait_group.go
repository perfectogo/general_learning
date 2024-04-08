package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go run1("hello, world", &wg)
	go run("bye", &wg)

	wg.Wait()
}

func run(s string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}

	wg.Done()

}

func run1(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}

	wg.Done()

}
