package routes

import (
	"github.com/darshan744/go-bookstore/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/book", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}
