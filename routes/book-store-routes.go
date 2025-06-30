// Package routes defines and registers all HTTP route handlers for the server.
// It connects the Gin routes to their corresponding controller functions.
package routes

import (
	"github.com/darshan744/go-bookstore/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API endpoints for the application using the given Gin engine.
//
// The following routes are available:
// - GET    /book           → Get all books
// - POST   /book           → Create a new book
// - GET    /book/:bookId   → Get a book by its ID
// - PUT    /book/:bookId   → Update a book by its ID
// - DELETE /book/:bookId   → Delete a book by its ID
func RegisterRoutes(router *gin.Engine) {
	router.GET("/book", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}
