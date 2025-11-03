package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Connect() {
	// Neon provides connection strings like:
	// postgres://user:pass@ep-cool-cloud-123456.us-east-1.aws.neon.tech/neondb
	connectionString := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to Neon DB:", err)
	}

	DB = db
	fmt.Println("✅ Connected to Neon DB!")
}
