package middlewares

import (
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		repositories.MakeLogEntry(c).Info("Incoming request...")
		return next(c)
	}
}