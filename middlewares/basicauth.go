package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// BasicAuth Method middleware for setting method
// TODO: this is mock, use db(or usermanagement layer) for auth
func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "blasphemy" && password == "thisissparta" {
			return true, nil
		}
		return false, nil
	})
}
