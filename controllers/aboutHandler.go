package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET "/about"
func AboutHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "about", nil)
}