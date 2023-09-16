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

### JSON schema
```
{
    "name": "string",
    "age": "number",
    "address": "string"
}
```