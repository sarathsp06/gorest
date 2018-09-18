package middlewares

import (
	"github.com/sarathsp06/gorest/utils"
	"github.com/labstack/echo"
)

// KeyRequestID key in echo context
const KeyRequestID = "RequestID"

// RequestID middleware for setting request id
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		c.Set(KeyRequestID, utils.GetUniqueID())
		return next(c)
	})
}
