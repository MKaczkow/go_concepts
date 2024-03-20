package services

import "web_app_bis/models"

type BookServiceInterface interface {
	CreateBook(*models.Book) error
	GetBook(*string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateBook(*models.Book) error
	DeleteBook(*string) error
}
