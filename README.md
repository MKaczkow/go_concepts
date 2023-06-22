# go_concepts
repo for basic tutorial-based Golang study  

### based on
[booking app tutorial](https://www.youtube.com/watch?v=yyUHQIec83I&t=3693s&ab_channel=TechWorldwithNana) (booking app + some basics)  
[basics tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org) (most basics)  
[web app tutorial](https://www.youtube.com/watch?v=LOn1GUsjOF4&ab_channel=DavidAlsh) (web app + Gin)  
[web app bis tutorial](https://www.youtube.com/watch?v=vDIAwtGU9LE&ab_channel=DevProblems) (web app + Gin)

### issues
* `func (u *UserService) CreateUser(user *models.User)` What does the (u *UserService) part mean? Likely, it specifies return type, it so?
* `func (u *UserService) CreateUser(user *models.User)` What does * mean? It looks like a pointer...
* `if err := ctx.ShouldBindJSON(&user); err != nil {` What about ShouldBindJSON?
* `query := bson.D{bson.E{Key: "name", Value: name}}` What is bson.D or bson.E? (Encode / Decode??)
* add tests -> web_app_bis??
