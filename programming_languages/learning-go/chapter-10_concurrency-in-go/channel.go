package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var number int
	var m sync.Mutex = sync.Mutex{}

	go func() {
		m.Lock()         // Mutexni qo'lga kiritamiz
		defer m.Unlock() // Mutexni qo'lga kirishni o'chiramiz

		number++ // number o'zgaruvchisini o'zgartiramiz

	}()

	//time.Sleep(1 * time.Second)

	m.Lock()         // Main gorutini mutexni qo'lga kiritadi
	defer m.Unlock() // Main gorutini mutexni qo'lga kirishni o'chiradi

	fmt.Println(number) // number o'zgaruvchisini chop etamiz
}
