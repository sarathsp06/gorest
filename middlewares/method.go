package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// KeyMethod stores the key used to store
const KeyMethod = "Method"

// Method middleware for setting method
// why? Some routers do not allow DELETE,PATCH etc requests
// in that scenarios POST request with the http method `x` set as _method query parameter would
// interpreted as having `x` http method
func Method(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		method := c.Request().Method
		customMethod := strings.ToUpper(c.QueryParam("_method"))
		switch customMethod {
		case http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut:
			method = customMethod
		}
		//TODO : Revisit the following
		//Why : Check if it is changing the method at the route
		{
			c.Request().Method = method
			c.SetRequest(c.Request())
		}
		c.Set(KeyMethod, method)
		return next(c)
	})
}
