// Handles configurations for the server
package configs

import (
	"fmt"
	"log"
	"os"
	"github.com/darshan744/go-bookstore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


// The db struct holder
// Using this we perform DB operations
var dbConnection *gorm.DB

// This function connects to a Postgres Database
// if db connection fails program exits
// It uses Gorm for ORM and DB Connection
// Gorm uses pgx package for postgres internally 
func Connect() {
	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)
	//conection string for postgres
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	db.AutoMigrate(&models.Book{})
	dbConnection = db
}
// Helper function for getting db struct
func GetDb() *gorm.DB {
	return dbConnection
}
