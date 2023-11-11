package database

import (
	"log"
	"os"

	"github.com/nathanaelpganata/go-fullstack-shortener/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		log.Println("connected to db")
	}

	DB = db

	if err := db.AutoMigrate(&models.Link{}); err != nil {
		log.Fatal("failed to migrate")
	}
}
