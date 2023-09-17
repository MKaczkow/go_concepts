package services

import (
  "context"
  "errors"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"

  "web_app_bis/models"
)

type BookService struct {
	bookcollection *mongo.Collection
	ctx context.Context
}

func NewBookService(collection *mongo.Collection, ctx context.Context) BookServiceInterface {
	return &BookService{
		bookcollection: collection,
		ctx: ctx,
	}
}

func (u *BookService) CreateBook(book *models.Book) error {
	_, err := u.bookcollection.InsertOne(u.ctx, book)
	return err
}

func (u *BookService) GetBook(title *string) (*models.Book, error) { 
	var book *models.Book
	query := bson.D{bson.E{Key: "title", Value: title}}
	err := u.bookcollection.FindOne(u.ctx, query).Decode(&book) 
	return book, err
}

func (u *BookService) GetAll() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err := u.bookcollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var book *models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(books) == 0 {
		return nil, errors.New("no books found")
	}
	return nil, nil
}

func (u *BookService) UpdateBook(book *models.Book) error {
	filter := bson.D{primitive.E{Key: "book_title", Value: book.Title}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: book.Title}, 
		primitive.E{Key: "age", Value: book.Author}, 
		primitive.E{Key: "year", Value: book.Year},
		primitive.E{Key: "abstract", Value: book.Abstract},
	}}}

	result, _ := u.bookcollection.UpdateOne(u.ctx, filter, update)

	if result.ModifiedCount != 1 {
		return errors.New("failed to update book")
	}
	return nil
}

func (u *BookService) DeleteBook(title *string) error {
	filter := bson.D{primitive.E{Key: "title", Value: title}}
	result, _ := u.bookcollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("failed to delete book")
	}
	return nil
}