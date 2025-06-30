// Package models defines the data models used by the application.
// It includes structs for Book representation, URI parameters, and update input.
package models

import (
	"gorm.io/gorm"
)

// Book represents a book record in the database.
// It embeds gorm.Model which provides ID, CreatedAt, UpdatedAt, and DeletedAt fields.
type Book struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name" binding:"required"` // Name is the title of the book
	Author      string `json:"author" binding:"required"`               // Author is the name of the author
	Publication string `json:"publication" binding:"required"`          // Publication is the publishing house
}

// BookUri represents URI parameters used when accessing a specific book by ID.
type BookUri struct {
	BookId string `uri:"bookId" binding:"required"` // BookId is the ID of the book in the route
}

// BookUpdate represents the JSON body for updating a book.
// Fields are optional and will only be updated if provided.
type BookUpdate struct {
	Name        string `gorm:"not null" json:"name"` // Updated name of the book
	Author      string `json:"author"`               // Updated author name
	Publication string `json:"publication"`          // Updated publication name
}
