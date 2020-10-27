package controller

import (
	"net/http"

	"github.com/xaoirse/xgo/graph/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func LoginPage(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	}
}

func Login(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		s := &model.Session{
			Username: c.FormValue("username"),
		}
		s.Save(c, db)
		return c.Redirect(http.StatusSeeOther, "/")
	}
}
