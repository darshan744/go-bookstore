package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbConnection *gorm.DB

func init() {
	fmt.Println("Config init")
}
func Connect() {

	var (
		host   = os.Getenv("DB_HOST")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)
	fmt.Println(host)
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	dbConnection = db
}

func GetDb() *gorm.DB {
	return dbConnection
}
