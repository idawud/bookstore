# Bookstore Api
This a simple api in Go goFiber Framework, GORM and swaggo

Libraries
- [goFiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/jinzhu/gorm)
- [fiber-swagger](https://github.com/arsmn/fiber-swagger)

Setup
Install the [swaggo](github.com/swaggo/swag) cmd with `go get -U` see documentations and ensure the GOPATH is available in the system PATH.
- Run `swag init` this will generate the go swagger specification in `./docs` dir.
- Run `go run main.go` [NB: Ensure all dependent libraries are already downloaded]

Run with Docker
- `docker build -t idawud/bookstore:v1 .`
- `docker run -it -p 3000:3000 idawud/bookstore:v1`