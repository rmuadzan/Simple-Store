package controllers

import (
	"net/http"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

// GET "logout"
func LogoutHandler(ctx echo.Context) error {
	repositories.DeleteCookie(ctx, "jwt")
	return ctx.Redirect(http.StatusSeeOther, "/login")
}