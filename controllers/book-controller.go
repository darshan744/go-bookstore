// Handles incoming Request for the server
package controllers

import (
	"fmt"
	"net/http"
	"github.com/darshan744/go-bookstore/configs"
	"github.com/darshan744/go-bookstore/models"
	"github.com/gin-gonic/gin"
)

// Route -> /book/
// Method POST
// Creates a Book and returns it
func CreateBook(context *gin.Context) {
	book := models.Book{}
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.String(http.StatusBadRequest, "Request Body is not Correct")
		return
	}
	db := configs.GetDb()
	tx := db.Begin()
	dBerr := tx.Create(&book).Error
	if dBerr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	tx.Commit()
	context.JSON(200, book)
}

// Route -> /book/
// Method -> GET
// Get All Books
func GetBook(context *gin.Context) {
	var books []models.Book
	db := configs.GetDb()
	err := db.Find(&books).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	fmt.Println(books)
	context.JSON(200, gin.H{"data": books})
}

// Route -> /book/{bookId}
// Method -> GET
// Returns the Book by its id
func GetBookById(context *gin.Context) {

	var uri models.BookUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.String(http.StatusBadRequest, "Invalid Request Body")
		return
	}
	var book models.Book
	db := configs.GetDb()
	db.First(&book, "id = ?", uri.BookId)

	context.JSON(http.StatusOK, gin.H{"book": book})
}

// Route -> /book/{bookId}
// Method -> DELETE
// Deletes the book by its id
func DeleteBook(context *gin.Context) {
	var uri models.BookUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println(uri)
	db := configs.GetDb()
	var book models.Book
	db.Delete(&book, uri.BookId)
	context.JSON(http.StatusOK, gin.H{"message": "deleted", "data": book})
}

// Route -> /book/{bookId}
// Method -> PUT
// Updates the Book based on Id
func UpdateBook(context *gin.Context) {
	var updateBook models.BookUpdate
	var uri models.BookUri

	if uriErr := context.ShouldBindUri(&uri); uriErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": uriErr})
		return
	}
	if bodyErr := context.ShouldBindJSON(&updateBook); bodyErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": bodyErr})
		return
	}
	db := configs.GetDb()
	var book models.Book
	if err := db.First(&book, uri.BookId).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := db.Model(&book).Updates(updateBook).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(200, gin.H{"book": book})
}
