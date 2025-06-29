package main

import (
	"github.com/darshan744/go-rest-server/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
func main() {
	router := gin.Default()
	database.Connect()
	router.Run()
}
