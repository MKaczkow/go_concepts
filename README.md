# go_concepts
repo for basic tutorial-based Golang study  

### based on
[booking app tutorial](https://www.youtube.com/watch?v=yyUHQIec83I&t=3693s&ab_channel=TechWorldwithNana) (booking app + some basics)  
[basics tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org) (most basics)  
[web app tutorial](https://www.youtube.com/watch?v=LOn1GUsjOF4&ab_channel=DavidAlsh) (web app + Gin)  
[web app bis tutorial](https://www.youtube.com/watch?v=vDIAwtGU9LE&ab_channel=DevProblems) (web app + Gin)

### issues
* (DONE) `func (u *UserService) CreateUser(user *models.User)` What does the (u *UserService) part mean? Likely, it specifies return type, it so? (`CreateUser(user *models.User)` part specifies the return type)
* (DONE) `func (u *UserService) CreateUser(user *models.User)` What does * mean? It looks like a pointer... (yep, it's a pointer)
* `if err := ctx.ShouldBindJSON(&user); err != nil {` What about ShouldBindJSON?
* (DONE) `query := bson.D{bson.E{Key: "name", Value: name}}` What is bson.D or bson.E? (these are two out of four basic types of BSON documents, which are the basic unit of data in MongoDB, ref: https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/)
* add tests -> web_app_bis??
* ```
  update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "name", Value: user.Name}, 
		primitive.E{Key: "age", Value: user.Age}, 
		primitive.E{Key: "address", Value: user.Address},
	}}}
    ```
    Why changging Bson.E -> primitive.E fixed erorrs??
* (DONE) difference between `=` and `:=` in Golang? (`:=` means 'declare and assign', while `=` means 'assign')
* (DONE) can you return empty values in Golang? (yes, you can return empty values in Golang, in fact it's quite often a prefered way of doing things - so `return nil` or just `return` and don't bother)

### notes
* golang function syntax:  
`func functionName(parameter1 type1, parameter2 type2, parameterN typeN) returnType {
	   //function body
}`