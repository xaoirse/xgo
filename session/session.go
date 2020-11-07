package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	// sessionStr represents a string of session key.
	sessionStr = "GSESSION"
)

// Secure is a middleware that check session for access
// Redirect to /login/ if session not found
func Secure(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, err := Get(c); err == nil || c.Path() == "/" {
			return handler(c)
		}
		return c.Redirect(http.StatusSeeOther, "/login/")
	}
}

// Get returns a session for the current request.
func Get(c echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(sessionStr, c)
	return sess, err
}

// Save saves the current session.
func Save(c echo.Context) error {
	sess, _ := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   20,
	}
	return saveSession(c, sess)
}

// Delete the current session.
func Delete(c echo.Context) error {
	sess, _ := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	return saveSession(c, sess)
}

func saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return nil
}

// SetValue sets a key and a value.
func SetValue(c echo.Context, key string, value interface{}) error {
	sess, _ := Get(c)
	sess.Values[key] = value
	return nil
}

// GetValue returns value of session.
func GetValue(c echo.Context, key string) interface{} {
	sess, _ := Get(c)
	if sess != nil {
		if v, ok := sess.Values[key]; ok {
			return v
		}
	}
	return nil
}
