// Package configs handles the database configuration and connection logic
// for the application. It establishes a connection to a PostgreSQL database
// using GORM and provides access to the DB instance.
package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/darshan744/go-bookstore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

// Connect establishes a connection to the PostgreSQL database using
// environment variables for configuration. It uses GORM as the ORM
// and automatically runs migrations for the Book model.
// If the connection fails, the application logs the error and exits.
func Connect() {
	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)

	connectionString := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		host, user, pass, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}

	// AutoMigrate will create the "books" table if it doesn't exist
	db.AutoMigrate(&models.Book{})

	dbConnection = db
}

// GetDb returns a pointer to the initialized *gorm.DB instance.
// This is used across the application to perform database operations.
func GetDb() *gorm.DB {
	return dbConnection
}
