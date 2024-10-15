package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// DB is a global variable to hold the database connection
var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	_, err = gorm.Open(mysql.Open("root:44125@tcp(127.0.0.1:3306)/nostalgia-be?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Println("Database connected successfully")
}

// Close disconnects the database connection
func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("Could not get database instance: %v", err)
	}
	err = db.Close()
	if err != nil {
		log.Fatalf("Could not close the database: %v", err)
	}
	log.Println("Database connection closed")
}
