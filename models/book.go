package models

import (
	"gorm.io/gorm"
)


type Book struct {
	// this has some default fields
	// so we can just use them 
	// instead of creating our own
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}
