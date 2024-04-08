package main

import (
	"fmt"
	"sync"
)

var counter int
var counterMutex sync.Mutex

func increment() {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	counter++
}

func main() {
	
	//var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		//wg.Add(1)
		
		go func() {
			//defer wg.Done()
			increment()
		}()
	}

	//wg.Wait()
	fmt.Println("Final counter value:", counter)
}
