package controllers

import ("github.com/gin-gonic/gin" 
	"github.com/darshan744/go-bookstore/models")

func CreateBook(context *gin.Context) {
	book := models.Book{}
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.String(404 , "Request Body is not Correct")
		
	}
	context.JSON(200 , book)
}

func GetBook(context *gin.Context) {}

func GetBookById (context *gin.Context) {}

func DeleteBook(context *gin.Context) {}

func UpdateBook (context *gin.Context) {}
