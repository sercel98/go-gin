package models

import "github.com/jinzhu/gorm"

//Struct Video represents a Video models with his attributes

type Video struct {
	//gorm.Model injects: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Title       string `json:"title" gorm:"size:255; not null"`
	Description string `json:"description" gorm:"size:400; not null"`
	URL         string `json:"url" gorm:"not null"`
}
