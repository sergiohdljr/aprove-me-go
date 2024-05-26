package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InnitDB() {
	var err error

	uri := "host=localhost user=postgres password=87575162 dbname=aprove_me port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("database connection estabilished")
}
