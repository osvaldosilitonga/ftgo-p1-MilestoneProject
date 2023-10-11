package handler

import (
	"fmt"
	"klepon/config"
	"log"
)

func GetOrders() {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	fmt.Println("Test Get Orders")
}
