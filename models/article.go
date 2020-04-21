package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Struct Article represents an Article models with his attributes
type Article struct {
	gorm.Model
	Title         string    `json:"title" gorm:"size:100; not null"`
	Author        string    `json:"author" gorm:"size:70; not null"`
	PublishedDate time.Time `json:"publishedDate" gorm:"type:time; not null"`
	Content       string    `json:"content" gorm:"size:600; not null"`
}
