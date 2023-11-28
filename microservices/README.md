# Microservices exercise with MongoDB, Golang (Gin) and Docker

### TODO
- [x] docker desktop works
- [x] basic, helloworld microservices work
- [x] docker-compose works
- [x] push to dockerhub
- [x] design reasonable microservices
    - database
    - database-api
    - omdb-api
    - recommendation-engine
- [ ] design APIs and DB schemas
    - recomendation-engine
    - which data about users to store
    - database-api (CRUD, so easy)
    - omdb-api

### Goals of this exercise
* to learn and have fun :smile:
* to make systems written in different languages work together with JSON API (or other API, but most likely JSON)
* to setup plain and simple docker-compose for the whole system
* to research basic recommendation engines/ algos/ models in Python (likely NLP or graph based)
* to setup automated testing/ building/ etc. using Github Actions

### Containers / microservices
* **database** (mongoDB) - just container with mongoDB
* **database-api** (Golang + Gin) - API to communicate with database
* **omdb-api** (Golang + Gin) [link](https://www.omdbapi.com/) - API to communicate with OMDB
* **recommendation-engine** (Python + FastAPI) - engine to recommend movies

### Example OMDB API calls
* http://www.omdbapi.com/?apikey={apikey}&t=Inception
* http://www.omdbapi.com/?apikey={apikey}&i=tt1285016

### Example internal API calls
* http://localhost:8080/movies?t=Inception
* http://localhost:8080/movies?i=tt1285016


### Usage
* `docker-compose up -d`
* `docker build . -t lorca19/golang-microservices:v0.2.0`
* `docker run -p 8080:8080 lorca19/golang-microservices:v0.2.0`

### Dockerhub 
* `docker login`
* `docker push lorca19/golang-microservices:v0.2.0`

### MongoDB notes
* `createCollection` vs `createView`
* most operations, are analog to SQL (e.g. `find`, `create`, `distinct`, `count`)
* Go Driver provides four main types for working with BSON data:
    - **D**: An ordered representation of a BSON document (slice)
    - **M**: An unordered representation of a BSON document (map)
    - **A**: An ordered representation of a BSON array
    - **E**: A single element inside a D type
    