package pghelper

import (
	"github.com/alexwitedja/packform-test-api/backend/models"
	"log"

	"github.com/jinzhu/gorm"
)

// ConnectDB Connect to pg db and return gorm.DB
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=root")

	if err != nil {

		log.Fatal("Failed to connect to pg db.")

	}

	log.Printf("Connected to pg db!")

	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.OrderItem{})

	return db
}
