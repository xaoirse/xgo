package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xaoirse/xgo/session"
)

type User struct{ Username string }

// Home controller
func GetHome(c echo.Context) error {
	username := session.GetValue(c, "username")
	var name string
	if username == nil {
		name = "XGO"
	} else {
		name = username.(string)
	}
	return c.Render(http.StatusOK, "home.html", User{Username: name})
}
