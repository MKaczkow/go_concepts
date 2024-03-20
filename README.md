# go_concepts
repo for basic tutorial-based Golang study  

### todo
- [ ] finish `book`
- [ ] smth with `Hanoi tower`
- [ ] finish LP (liear programming) task  
- [x] fix `web_app_bis` (if no JSON is provided, some routes crash)
- [x] init `book`
- [x] clean `issues` section (for now)
- [x] fix github actions
- [x] get book "Język Go. Tworzenie idiomatycznego kodu w praktyce" [example link](https://ulubionykiosk.pl/wydawnictwo/jezyk-go-tworzenie-idiomatycznego-kodu-w-praktyce?gclid=CjwKCAiAyp-sBhBSEiwAWWzTnidWyVtzQT6rU82MAzZSNY6u-Vx3KuyetmuLR5GSGNId6kPF5nr_IxoCv5AQAvD_BwE) [other link](https://helion.pl/ksiazki/jezyk-go-tworzenie-idiomatycznego-kodu-w-praktyce-jon-bodner,jegotw.htm#format/e)
- [x] finish `web_app_bis` tutorial (mostly done, need to check if everything works OK)
~~ [ ] work on new task [`microservices`](./microservices/)~~

### basic usage
* cd `app_name`
* `go build -v ./`
* `go test -v ./`
* `go run main.go`

### based on
[booking app tutorial](https://www.youtube.com/watch?v=yyUHQIec83I&t=3693s&ab_channel=TechWorldwithNana) (booking app + some basics)  
[basics tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org) (most basics)  
[web app tutorial](https://www.youtube.com/watch?v=LOn1GUsjOF4&ab_channel=DavidAlsh) (web app + Gin)  
[web app bis tutorial](https://www.youtube.com/watch?v=vDIAwtGU9LE&ab_channel=DevProblems) (web app + Gin)

### issues
* How should modules / packages be organized? How are they organized in real-life large projects? [docs](https://golang.org/doc/code.html#Organization)
* How specifically, does `hash maps` work in Golang? 
* (DONE) `func (u *UserService) CreateUser(user *models.User)` What does the (u *UserService) part mean? Likely, it specifies return type, it so? (`CreateUser(user *models.User)` part specifies the return type)
* (DONE) `func (u *UserService) CreateUser(user *models.User)` What does * mean? It looks like a pointer... (yep, it's a pointer)
* (DONE) `if err := ctx.ShouldBindJSON(&user); err != nil {...}` What about ShouldBindJSON? (the syntax means, that we first create err variable, assign value to it and only then check if function returned any errors)
* (DONE) `query := bson.D{bson.E{Key: "name", Value: name}}` What is bson.D or bson.E? (these are two out of four basic types of BSON documents, which are the basic unit of data in MongoDB, ref: https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/)
* ```
  update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "name", Value: user.Name}, 
		primitive.E{Key: "age", Value: user.Age}, 
		primitive.E{Key: "address", Value: user.Address},
	}}}
    ```
    Why changging Bson.E -> primitive.E fixed erorrs ?? (from https://pkg.go.dev/go.mongodb.org/mongo-driver/bson:  
	**M** - unordered representation of a BSON document  
	**D** - ordered representation of a BSON document  
	**E** - a BSON element for a D  
	)
* (DONE) difference between `=` and `:=` in Golang? (`:=` means 'declare and assign', while `=` means 'assign')
* (DONE) can you return empty values in Golang? (yes, you can return empty values in Golang, in fact it's quite often a prefered way of doing things - so `return nil` or just `return` and don't bother)

### notes
* golang function syntax:  
`func functionName(parameter1 type1, parameter2 type2, parameterN typeN) returnType {
	   //function body
}`
