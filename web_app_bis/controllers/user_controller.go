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
	return ctx.JSON(http.StatusOK, "")
}

func (u *UserController) GetUser(ctx *gin.Context) (*models.User, error) {
	return ctx.JSON(http.StatusOK, "")
}

func (u *UserController) GetAll(ctx *gin.Context) ([]*models.User, error) {
	return ctx.JSON(http.StatusOK, "")
}

func (u *UserController) UpdateUser(ctx *gin.Context) error {
	return ctx.JSON(http.StatusOK, "")
}

func (u *UserController) DeleteUser(ctx *gin.Context) error {
	return ctx.JSON(http.StatusOK, "")
}