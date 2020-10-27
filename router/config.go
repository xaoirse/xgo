package router

import (
	"io"
	"log"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB

// Template for echo
type Template struct {
	templates *template.Template
}

// Render for echo
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Start Echo
func Start(db *gorm.DB, port *string, secret []byte) {
	if len(secret) == 0 {
		secret = securecookie.GenerateRandomKey(32)
		if secret == nil {
			log.Fatal("Failed to generate random secret!\nPlease insert a secret like this: -s [secret]")
		}
	}

	e := New(db, secret)

	t := &Template{templates: template.Must(template.ParseGlob("template/*.html"))}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":" + *port))
}
