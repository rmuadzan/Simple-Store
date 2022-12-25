package controllers

import (
	"net/http"
	"simple-catalog-v2/models"

	"github.com/labstack/echo/v4"
)

func LogoutHandler(ctx echo.Context) error {
	models.DeleteCookie(ctx, "jwt")
	return ctx.Redirect(http.StatusSeeOther, "/login")
}