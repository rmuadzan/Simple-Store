package middlewares

import (
	"simple-catalog-v2/models"

	"github.com/labstack/echo/v4"
)

func MiddlewareContextValue(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx = &models.ContextValue{ctx}
		return next(ctx)
	}
}