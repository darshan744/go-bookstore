package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darshan744/go-bookstore/configs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	fmt.Println("Main init")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
func main() {
	router := gin.Default()
	configs.Connect()
	router.Run()
}
