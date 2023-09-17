// models_test.go

package models

import (
	"testing"
)

func TestBook(t *testing.T) {
	// Create a sample Book
	book := Book{
		Title:    "Sample Book",
		Author:   "John Author",
		Year:     2022,
		Abstract: "This is a sample book abstract.",
	}

	// Test the fields of the Book struct
	if book.Title != "Sample Book" {
		t.Errorf("Expected Title to be 'Sample Book', but got %s", book.Title)
	}

	if book.Author != "John Author" {
		t.Errorf("Expected Author to be 'John Author', but got %s", book.Author)
	}

	if book.Year != 2022 {
		t.Errorf("Expected Year to be 2022, but got %d", book.Year)
	}

	if book.Abstract != "This is a sample book abstract." {
		t.Errorf("Expected Abstract to be 'This is a sample book abstract.', but got %s", book.Abstract)
	}
}
