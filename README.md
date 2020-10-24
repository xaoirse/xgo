## Golang web app template

Simple web app for using GraphQL, GORM (with sqlite3) and Echo together. Just clone and run it.

## Requirements
`go.mod` has all requirements but if you want to install manually:
- `git clone github.com/xaoirse/xgo` for clone XGO
- `cd XGO`
- `go get github.com/99designs/gqlgen` for get gqlgen.<br/> 
- `go get -u gorm.io/gorm` for get gorm.<br/> 
- `go get -u github.com/labstack/echo/...` for get echo.<br/> 



## Usage
- Define your models in `graph/schema.graphqls`
- `go run github.com/99designs/gqlgen generate`
- Insert your models in `graph/model/model.go`
- Complate resolvers in `schema.resolvers.go`
- Routing in `router/router.go`
- Controllers are in `controller` or anywhere
- `go run server.go` or `go build`

*NOTE: Find and replace mod name if you want*

## TODO
- [ ] go:generate (Maybe not)
- [ ] general router.go for using any router in config.go (Maybe not)