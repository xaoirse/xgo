package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xaoirse/xgo/session"
)

func GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func PostLogin(c echo.Context) error {
	session.SetValue(c, "username", c.FormValue("username"))
	session.Save(c)
	return c.Redirect(http.StatusSeeOther, "/")
}
