# Go Concepts
Repo for basic tutorial-based Golang study  

[![Housekeeping CI](https://github.com/MKaczkow/go_concepts/actions/workflows/housekeeping-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/housekeeping-ci.yml)  
[![Learning Go Testing CI](https://github.com/MKaczkow/go_concepts/actions/workflows/learning-go-testing-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/learning-go-testing-ci.yml)  
[![Web App Bis CI](https://github.com/MKaczkow/go_concepts/actions/workflows/web-app-bis-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/web-app-bis-ci.yml)  
[![Web Crawler CI](https://github.com/MKaczkow/go_concepts/actions/workflows/web-crawler-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/web-crawler-ci.yml)  
[![Regex Engine CI](https://github.com/MKaczkow/go_concepts/actions/workflows/regex-engine-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/regex-engine-ci.yml)  
[![Gilded Rose CI](https://github.com/MKaczkow/go_concepts/actions/workflows/gilded-rose-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/gilded-rose-ci.yml)  
[![Monkey Interpreter CI](https://github.com/MKaczkow/go_concepts/actions/workflows/monkey-interpreter-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/monkey-interpreter-ci.yml)  
[![Monkey Compiler CI](https://github.com/MKaczkow/go_concepts/actions/workflows/monkey-compiler-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/monkey-compiler-ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-76.1%25-brightgreen)

---

### todo
- [x] dubious shadowing
- [ ] monkey
	- [ ] compiler
	- [x] macros
	- [x] interpreter
- [ ] regex engine
- [ ] small cleanup
- [x] only trigger CI pipeline when changes occur in given directory
- [x] do `gilded rose` kata
- [x] web crawler with `colly`
  - [x] mop up
  - [x] choose storage method (like Elasticsearch, MongoDB, etc.)
  - [x] upgrade crawler to store data in chosen storage
  - [x] crawl all places on target website
  - [x] design data model (i.e. what info about the place should be stored)
  - [x] upgrade crawler to crawl all requested data
- [x] smth with `Hanoi tower`
- [x] finish `book`
- [x] fix `web_app_bis` (if no JSON is provided, some routes crash)
- [x] init `book`
- [x] clean `issues` section (for now)
- [x] fix github actions
- [x] get book "Język Go. Tworzenie idiomatycznego kodu w praktyce" [example link](https://ulubionykiosk.pl/wydawnictwo/jezyk-go-tworzenie-idiomatycznego-kodu-w-praktyce?gclid=CjwKCAiAyp-sBhBSEiwAWWzTnidWyVtzQT6rU82MAzZSNY6u-Vx3KuyetmuLR5GSGNId6kPF5nr_IxoCv5AQAvD_BwE) [other link](https://helion.pl/ksiazki/jezyk-go-tworzenie-idiomatycznego-kodu-w-praktyce-jon-bodner,jegotw.htm#format/e)
- [x] finish `web_app_bis` tutorial (mostly done, need to check if everything works OK)  

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
```go
func functionName(parameter1 type1, parameter2 type2, parameterN typeN) returnType {
	   //function body
}
```
* golang comparision of objects also takes into account the actual data location in memory, so two objects with the same values, but different memory locations are not equal (example from *Writing an Interpreter in Go*):
```go
name1 := &object.String{Value: "name"}
monkey := &object.String{Value: "Monkey"}
pairs := map[object.Object]object.Object{}

pairs[name1] = monkey
fmt.Printf("pairs[name1]=%+v\n", pairs[name1])
// => pairs[name1]=&{Value:Monkey}

name2 := &object.String{Value: "name"}
fmt.Printf("pairs[name2]=%+v\n", pairs[name2])
// => pairs[name2]=<nil>
fmt.Printf("(name1 == name2)=%t\n", name1 == name2)
// => (name1 == name2)=false
```

### cool resources
* `most important basics`
	* [Go cheatsheet](https://devhints.io/go)
	* [Google style guide for Golang](https://google.github.io/styleguide/go/)
	* [language specification](https://go.dev/ref/spec)
* `pointers, performance`
	* [stacks and pointers in Go](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
	* [pointers efficiency](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/)
* `on type parameters`
	* [Go 1.18 reease notes](https://tip.golang.org/doc/go1.18)
	* [2013 rejected proposal](https://go.googlesource.com/proposal/+/master/design/15292/2013-12-type-params.md)
	* [2021 accepted proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
	* [summary of discussion](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/view#heading=h.q325c8t1k569)
* `frameworks`
	* [top 5 frameworks geeksforgeeks](https://www.geeksforgeeks.org/top-5-golang-frameworks-in-2020/)
	* [web frameworks medium](https://medium.com/@livajorge7/exploring-the-best-golang-web-frameworks-a-comprehensive-guide-to-building-web-applications-with-daa3ae52b15c)
	* [another post on cool, best, can't miss frameworks](https://www.bacancytechnology.com/blog/golang-web-frameworks)
* `gin`
	* [gin framework blogpost](https://www.tabnine.com/blog/golang-gin/)
	* [gin tutorial yt](https://www.youtube.com/playlist?list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w)
	* [git tutorial yt](https://www.youtube.com/watch?v=vDIAwtGU9LE&ab_channel=DevProblems)
	* [another gin tutorial yt](https://www.youtube.com/watch?v=LOn1GUsjOF4&ab_channel=DavidAlsh)
	* [example REST API using gin](https://github.com/restuwahyu13/go-rest-api)
	* [another example REST API using gin](https://github.com/gothinkster/golang-gin-realworld-example-app)
* `colly`
	* [github examples](https://github.com/gocolly/colly/tree/master/_examples)
	* [some tutorial](https://dev.to/claudbytes/build-a-web-scraper-with-go-3jod)
	* [save data to csv after scrapping](https://webscraping.ai/faq/colly/how-do-i-save-the-scraped-data-to-a-file-using-colly)
* `misc`
	* [blogpost on bulldogjob](https://bulldogjob.pl/readme/pisz-w-jezyku-go-jak-senior)
	* [official docker image](https://hub.docker.com/_/golang)
	* [Athens alternative package sever](https://docs.gomods.io/)
	* [arxiv paper on usage of 'unsafe'](https://arxiv.org/pdf/2006.09973.pdf)
	* [effective go](https://go.dev/doc/effective_go) 
	* [A* implementation in Go](https://gist.github.com/egonelbre/10578266)
* `gilded rose`
	* resources in respective [README](./gilded_rose/README.md)
* `style guides`
	* google [style guide](https://google.github.io/styleguide/go/)
 	* uber [style guide](https://github.com/uber-go/guide/blob/master/style.md)
	
