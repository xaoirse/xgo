package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xaoirse/xgo/session"
)

func LoginPage(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	}
}

func Login(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		session.SetValue(c, "username", c.FormValue("username"))
		session.Save(c)
		return c.Redirect(http.StatusSeeOther, "/")
	}
}
