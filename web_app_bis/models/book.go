package models


type Book struct {
	Title string `json:"title" bson:"book_title"`
	Author string `json:"author" bson:"book_author"`
	Year int `json:"year" bson:"book_year"`
	Abstract string `json:"abstract" bson:"book_abstract"`
}
