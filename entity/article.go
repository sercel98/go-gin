package entity

import "time"

//Struct Article represents an Article entity with his attributes
type Article struct {
	Title         string
	Author        string
	PublishedDate time.Time
	Content       string
}