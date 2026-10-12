# Example web app

### Basis Setup
* (in terminal) 
  * *add ngrok authtoken if not added*
  * `ngrok http --url=well-closely-mutt.ngrok-free.app 8080`
* (in WSL) 
  * `cd web_app`
  * `go run main.go`
* (verify)
  * visit or curl GET `well-closely-mutt.ngrok-free.app/ping`

### Testing
`go test -cover ./...`

### Ngrok Reference
* [docs](https://dashboard.ngrok.com/get-started/setup/windows)