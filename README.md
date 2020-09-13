## Golang web app Template

Simple web app for using GraphQL, GORM and Echo together. 

## Requirements
- `git clone github.com/xaoirse/xgo` for clone XGO
- `cd XGO`
- `go get github.com/99designs/gqlgen` for install gqlgen.<br/> 
- `go get -u gorm.io/gorm` for install gorm.<br/> 
- `go get -u github.com/labstack/echo/...` for install echo.<br/> 


## Usage
- Define your models in `graph/schema.graphqls`
- `go run github.com/99designs/gqlgen generate`
- Insert your models in `graph/model/db.go`
- Complate resolvers in `schema.resolvers.go`
- Routin in `router/router.go`
- Controllers are in `controller`
- `go run server.go`

*NOTE: Find and replace mod name if you want*