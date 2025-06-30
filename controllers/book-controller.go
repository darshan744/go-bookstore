// Package controllers handles all incoming HTTP requests for the server.
// It defines the logic for creating, retrieving, updating, and deleting books.
package controllers

import (
	"github.com/darshan744/go-bookstore/configs"
	"github.com/darshan744/go-bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBook handles POST /book.
// It binds the incoming JSON to a Book model, saves it to the database,
// and returns the created book in the response.
func CreateBook(context *gin.Context) {
	book := models.Book{}
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.String(http.StatusBadRequest, "Request body is not correct")
		return
	}

	db := configs.GetDb()
	tx := db.Begin()
	if dbErr := tx.Create(&book).Error; dbErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": dbErr})
		tx.Rollback()
		return
	}
	tx.Commit()

	context.JSON(http.StatusOK, book)
}

// GetBook handles GET /book.
// It retrieves all books from the database and returns them as a JSON array.
func GetBook(context *gin.Context) {
	var books []models.Book
	db := configs.GetDb()

	if err := db.Find(&books).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBookById handles GET /book/:bookId.
// It extracts the book ID from the URI and returns the corresponding book.
func GetBookById(context *gin.Context) {
	var uri models.BookUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.String(http.StatusBadRequest, "Invalid request URI")
		return
	}

	var book models.Book
	db := configs.GetDb()
	if err := db.First(&book, "id = ?", uri.BookId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"book": book})
}

// DeleteBook handles DELETE /book/:bookId.
// It deletes the book with the given ID from the database.
func DeleteBook(context *gin.Context) {
	var uri models.BookUri
	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := configs.GetDb()
	var book models.Book
	if err := db.First(&book, uri.BookId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Delete(&book)
	context.JSON(http.StatusOK, gin.H{"message": "Deleted", "data": book})
}

// UpdateBook handles PUT /book/:bookId.
// It updates the book with the given ID using the fields provided in the request body.
func UpdateBook(context *gin.Context) {
	var updateBook models.BookUpdate
	var uri models.BookUri

	if err := context.ShouldBindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := context.ShouldBindJSON(&updateBook); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := configs.GetDb()
	var book models.Book
	if err := db.First(&book, uri.BookId).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := db.Model(&book).Updates(updateBook).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"book": book})
}
