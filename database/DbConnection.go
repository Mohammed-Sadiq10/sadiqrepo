package database

import (
	"fmt"
	"log"
	"os"
	"sadiq/Go_Rest_API/models"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
)

func DbConection() *gorm.DB{
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database configuration
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Establish database connection
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	
	// Auto-migrate the User model
	db.AutoMigrate(&models.User{})

	return db
}
