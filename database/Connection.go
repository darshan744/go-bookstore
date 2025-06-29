package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")	
)
var dbConnection *gorm.DB
func Connect() {
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v" , host , user , pass , dbname , port)
	db , err := gorm.Open(postgres.Open(connectionString) , &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	dbConnection = db
}
