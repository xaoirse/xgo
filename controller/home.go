package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Home controller
func Home(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", nil)
	}
}
