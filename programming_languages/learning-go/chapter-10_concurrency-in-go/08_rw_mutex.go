package main

import (
	"fmt"
	"sync"
)

type Database struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]string),
	}
}

func (db *Database) Read(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.data[key]
}

func (db *Database) Write(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func main() {
	db := NewDatabase()

	// Writing data concurrently
	for i := 0; i < 10; i++ {
		go func(n int) {
			db.Write(fmt.Sprintf("key%d", n), fmt.Sprintf("value%d", n))
		}(i)
	}

	// Reading data concurrently
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println("Reading:", db.Read(fmt.Sprintf("key%d", n)))
		}(i)
	}

	// Wait for goroutines to finish
}
