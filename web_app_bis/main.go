package main

import (
  "fmt"
  "context"
  "log"

  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options" 
  "go.mongodb.org/mongo-driver/mongo/readpref" 

  "web_app_bis/controllers"
  "web_app_bis/services"
)

var (
	ctx context.Context
	server *gin.Engine
	
	userservice services.UserServiceInterface
	usercontroller controllers.UserController
	usercollection *mongo.Collection

	bookservice services.BookServiceInterface
	bookcontroller controllers.BookController
	bookcollection *mongo.Collection

	mongoclient *mongo.Client
	err error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.NewUserController(userservice)
	
	bookservice = services.NewBookService(bookcollection, ctx)
	bookcontroller = controllers.NewBookController(bookservice)
	
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/api")
	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}