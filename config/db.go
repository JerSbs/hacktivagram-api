package config

import (
	"log"
	"os"
	"p2-livecode-3-JerSbs/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// Gunakan DATABASE_URL dari .env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	// Koneksi ke PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate semua model
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.UserActivityLog{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	log.Println("Connected to database and migrated successfully.")
}

func GetDB() *gorm.DB {
	return DB
}
