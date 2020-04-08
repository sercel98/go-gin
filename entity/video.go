package entity

//Struct Video represents a Video entity with his attributes
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
