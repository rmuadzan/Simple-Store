package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET "/login"
func LoginHandler(ctx echo.Context) error {
	if _, err := ctx.Cookie("jwt"); err == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	}
	
	return ctx.Render(http.StatusOK, "login", nil)
}