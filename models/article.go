package models

import "time"

//Struct Article represents an Article models with his attributes
type Article struct {
	Title         string
	Author        string
	PublishedDate time.Time
	Content       string
}
