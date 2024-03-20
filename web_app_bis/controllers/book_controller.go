package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"web_app_bis/models"
	"web_app_bis/services"
)

type BookController struct {
	BookService services.BookServiceInterface
}

func NewBookController(bookservice services.BookServiceInterface) BookController {
	return BookController{
		BookService: bookservice,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Exit the function
	}
	err := bc.BookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bc *BookController) GetBook(ctx *gin.Context) {
	booktitle := ctx.Param("title")
	book, err := bc.BookService.GetBook(&booktitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, book)
}

func (bc *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := bc.BookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, books)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Exit the function
	}
	err := bc.BookService.UpdateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bc *BookController) DeleteBook(ctx *gin.Context) {
	booktitle := ctx.Param("title")
	err := bc.BookService.DeleteBook(&booktitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bc *BookController) RegisterBookRoutes(rg *gin.RouterGroup) {
	bookroute := rg.Group("/books")
	bookroute.POST("/", bc.CreateBook)
	bookroute.GET("/:title", bc.GetBook)
	bookroute.GET("/", bc.GetAllBooks)
	bookroute.PUT("/:title", bc.UpdateBook)
	bookroute.DELETE("/:title", bc.DeleteBook)
}
