package model

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Session is a type for session
type Session struct {
	gorm.Model
	Username string
	Token    string
	Dest     string
	Age      int
}

// SessionChecker is a middleware that check session in db
// Redirect to /login/ if session not found
func SessionChecker(db *gorm.DB) func(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if IsSessionValid(c, db) || c.Path() == "/" {
				return handler(c)
			}
			return c.Redirect(http.StatusSeeOther, "/login/")
		}
	}
}

// IsSessionValid retruns true if session is in db
func IsSessionValid(c echo.Context, db *gorm.DB) bool {
	sess := GetSession(c)
	if sess != nil {
		var count int
		db.Find(sess).Count(&count)
		if count == 1 {
			return true
		}
	}
	return false
}

// GetSession return session if exists in echo.Context or nil
func GetSession(c echo.Context) *Session {
	sess, _ := session.Get("session", c)
	if sess.Values["token"] != nil &&
		sess.Values["username"] != nil {
		return &Session{
			Username: sess.Values["username"].(string),
			Token:    sess.Values["token"].(string),
			Age:      sess.Options.MaxAge,
		}
	}
	return nil
}

// Save a session in respose and db
func (s *Session) Save(c echo.Context, db *gorm.DB) error {
	sess, _ := session.Get("session", c)
	sess.Values["token"] = s.Token
	sess.Values["username"] = s.Username
	sess.Values["age"] = s.Age
	sess.Options.MaxAge = 5
	sess.Save(c.Request(), c.Response())

	return db.Create(s).Error
}
