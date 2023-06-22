package services

import "go.mongodb.org/mongo-driver/mongo"
import "web_app_bis/models"

type UserService struct {
	usercollection *mongo.Collection
	ctx context.Context
}

func NewUserService(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserService{
		usercollection: collection,
		ctx: ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	return nil
}

func (u *UserService) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(*models.User) error {
	return nil
}

func (u *UserService) DeleteUser(name *string) error {
	return nil
}