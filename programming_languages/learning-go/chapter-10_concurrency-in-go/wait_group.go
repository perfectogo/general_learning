package main

import (
	"fmt"
	"sync"
)

func main() {
	number := 0

	wg := &sync.WaitGroup{}
	var m = &sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				m.Lock()
				number++
				m.Unlock()
			}
		}(wg)
	}
	wg.Wait()
	fmt.Println(number)
}
