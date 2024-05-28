package database

import (
	"log"
	"os"
	"time"

	"github.com/sergiohdljr/aprove-me-go/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var NewLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  true,
	},
)

var db *gorm.DB

func InnitDB() {
	var err error

	uri := "host=localhost user=postgres password=87575162 dbname=aprove_me port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: NewLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	migrationErr := db.AutoMigrate(&models.Assignor{}, &models.Payment{})

	if migrationErr != nil {
		log.Fatalf("migration fails: %v", migrationErr)
	}

	log.Println("database connection estabilished")
}
