package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"web_app_bis/services"
	"web_app_bis/models"
)

type BookController struct {
	BookService services.BookServiceInterface
}

func NewBookController(bookservice services.BookServiceInterface) BookController {
	return BookController{
		BookService: bookservice,
	}
}

func (uc *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Exit the function
	}
	err := uc.BookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BookController) GetBook(ctx *gin.Context) {
	booktitle := ctx.Param("title")
	book, err := uc.BookService.GetBook(&booktitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, book)
}

func (uc *BookController) GetAll(ctx *gin.Context) {
	books, err := uc.BookService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, books)
}

func (uc *BookController) UpdateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // Exit the function
	}
	err := uc.BookService.UpdateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BookController) DeleteBook(ctx *gin.Context) {
	booktitle := ctx.Param("title")
	err := uc.BookService.DeleteBook(&booktitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // Exit the function
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BookController) RegisterBookRoutes(rg *gin.RouterGroup) {
	bookroute := rg.Group("/books")
	bookroute.POST("/", uc.CreateBook)
	bookroute.GET("/:title", uc.GetBook)
	bookroute.GET("/", uc.GetAll)
	bookroute.PUT("/:title", uc.UpdateBook)
	bookroute.DELETE("/:title", uc.DeleteBook)
}
