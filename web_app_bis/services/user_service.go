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
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserService) GetUser(name *string) (*models.User, error) { 
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user) 
	return nil, nil
}

func (u *UserService) GetAll() ([]*models.User, error) {
	users, 
	return nil, nil
}

func (u *UserService) UpdateUser(*models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "user_name", Value: user.Name}, 
		bson.E{Key: "user_age", Value: user.Age}, 
		bson.E{Key: "user_address", Value: user.Address}
	}}}
	result, _ = u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.ModifiedCount != 1 {
		return errors.New("failed to update user")
	}
	return nil
}

func (u *UserService) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	result, _ = u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("failed to delete user")
	}
	return nil
}