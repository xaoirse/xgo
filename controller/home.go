package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/xaoirse/xgo/graph/model"
)

// Home controller
func Home(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		u := model.GetSession(c)
		if u == nil {
			u = &model.Session{
				Username: "XGO",
			}
		}
		return c.Render(http.StatusOK, "home.html", u)
	}
}
