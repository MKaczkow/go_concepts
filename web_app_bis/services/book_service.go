package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"web_app_bis/models"
)

type BookService struct {
	bookcollection *mongo.Collection
	ctx            context.Context
}

func NewBookService(collection *mongo.Collection, ctx context.Context) BookServiceInterface {
	return &BookService{
		bookcollection: collection,
		ctx:            ctx,
	}
}

func (b *BookService) CreateBook(book *models.Book) error {
	_, err := b.bookcollection.InsertOne(b.ctx, book)
	return err
}

func (b *BookService) GetBook(title *string) (*models.Book, error) {
	var book *models.Book
	query := bson.D{bson.E{Key: "title", Value: title}}
	err := b.bookcollection.FindOne(b.ctx, query).Decode(&book)
	return book, err
}

func (b *BookService) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err := b.bookcollection.Find(b.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(b.ctx) {
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
	cursor.Close(b.ctx)
	if len(books) == 0 {
		return nil, errors.New("no books found")
	}
	return nil, nil
}

func (b *BookService) UpdateBook(book *models.Book) error {
	filter := bson.D{primitive.E{Key: "book_title", Value: book.Title}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: book.Title},
		primitive.E{Key: "age", Value: book.Author},
		primitive.E{Key: "year", Value: book.Year},
		primitive.E{Key: "abstract", Value: book.Abstract},
	}}}

	result, _ := b.bookcollection.UpdateOne(b.ctx, filter, update)

	if result.ModifiedCount != 1 {
		return errors.New("failed to update book")
	}
	return nil
}

func (b *BookService) DeleteBook(title *string) error {
	filter := bson.D{primitive.E{Key: "title", Value: title}}
	result, _ := b.bookcollection.DeleteOne(b.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("failed to delete book")
	}
	return nil
}
