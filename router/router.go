package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xaoirse/xgo/controller"
	"github.com/xaoirse/xgo/graph"
	"github.com/xaoirse/xgo/graph/generated"
	"github.com/xaoirse/xgo/repository"
	mySession "github.com/xaoirse/xgo/session"
)

// New return a new *Echo
func New(db *gorm.DB, secret []byte) *echo.Echo {

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Rep: repository.GetRepository()}}))

	e := echo.New()

	// TODO random secret generator
	// Note: Don't store your key in your source code. Pass it via an
	// environmental variable, or flag (or both), and don't accidentally commit it
	// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
	// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.

	e.Use(session.Middleware(sessions.NewCookieStore(secret)))
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Secure())

	// TODO uncomment for release or write better one
	// e.Use(middleware.Logger())

	// gqlgen
	e.GET("/playground/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query/", echo.WrapHandler(srv))

	// TODO a middleware for flood check

	// Root
	e.POST("/login/", controller.PostLogin)
	e.GET("/login/", controller.GetLogin)
	e.Use(mySession.Secure)
	e.GET("/", controller.GetHome)

	return e
}
