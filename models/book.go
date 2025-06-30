package models

import (
	"gorm.io/gorm"
)

type Book struct {
	// this has some default fields
	// so we can just use them
	// instead of creating our own
	gorm.Model
	Name        string `gorm:"not null" json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publication string `json:"publication" binding:"required"`
}

type BookUri struct {
	BookId string `uri:"bookId" binding:"required"`
}
type BookUpdate struct {
	
	Name        string `gorm:"not null" json:"name"`
	Author      string `json:"author"` 
	Publication string `json:"publication"` 
}
