# Microservices exercise with MongoDB, Golang (Gin) and Docker

### TODO
[x] docker desktop works
[x] basic, helloworld microservices work
[x] docker-compose works
[x] push to dockerhub
[] design reasonable microservices

### Containers
* database (mongoDB)
* api (Golang + Gin)

### Usage
* `docker-compose up -d`
* `docker build . -t lorca19/golang-microservices:v0.2.0`
* `docker run -p 8080:8080 lorca19/golang-microservices:v0.2.0`

### Dockerhub 
* `docker login`
* `docker push lorca19/golang-microservices:v0.2.0`
