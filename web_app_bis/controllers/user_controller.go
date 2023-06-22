package controllers

import "web_app_bis/services"

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(service services.UserServiceInterface) UserController {
	return &UserController{
		UserService: service,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) error {
	var usr models.User 
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (u *UserController) GetUser(ctx *gin.Context) (*models.User, error) {
	username := ctx.Param("name")
	user, err = uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) GetAll(ctx *gin.Context) ([]*models.User, error) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) UpdateUser(ctx *gin.Context) error {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (u *UserController) DeleteUser(ctx *gin.Context) error {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	rg.POST("/", uc.CreateUser)
	rg.GET("/:name", uc.GetUser)
	rg.GET("/", uc.GetAll)
	rg.PUT("/:name", uc.UpdateUser)
	rg.DELETE("/:name", uc.DeleteUser)
}