# Example web app bis

### Usage
``  
`go install`  
`go build`  
`go run main.go`  

### Testing
`go test -cover ./...`

### Routes
(for thunder client of postman)  
`GET` http://localhost:8080/api/users   
`GET` http://localhost:8080/api/users/:id  
`POST` http://localhost:8080/api/users  
`PUT` http://localhost:8080/api/users/:id  
`DELETE` http://localhost:8080/api/users/:id  
`GET` http://localhost:8080/api/books   
`GET` http://localhost:8080/api/books/:id  
`POST` http://localhost:8080/api/books  
`PUT` http://localhost:8080/api/books/:id  
`DELETE` http://localhost:8080/api/books/:id  

### JSON schema
* **user**
```
{
    "name": "string",
    "age": "number",
    "address": 
    {
        "street": "string",
        "city": "string",
        "state": "string"
    }
}
```
* **book**
```
{
    "title": "string",
    "author": "string",
    "year": "number"
    "abstract": "string"
}
```