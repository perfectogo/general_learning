package main

import (
	"examples/databases/config"
	"examples/databases/pkg/db"
	"fmt"
	"log"
)

func main() {

	config := config.LoadConfig()

	sqlDB, err := db.GetClickHouseConnection(config)

	if err != nil {
		log.Println("Failed to connect to ClickHouse:", err)
		return
	}

	log.Println("Connected to ClickHouse!")
	fmt.Println(sqlDB)

}
