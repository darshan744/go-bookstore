package main

import (
	"log"
	"os"

	"github.com/darshan744/go-bookstore/configs"
	"github.com/darshan744/go-bookstore/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
func main() {
	router := gin.Default()
	configs.Connect()
	routes.RegisterRoutes(router)
	router.Run()
}
