package controllers

import (
	"net/http"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

// GET "/"
func IndexHandler(ctx echo.Context) error {
	data := struct{
		IsLogged bool
		Fullname string
		Status string
	}{}

	userInfo := repositories.GetUserClaimsFromContext(ctx)
	data.IsLogged = true
	data.Fullname = userInfo.FullName
	data.Status = userInfo.Status

	return ctx.Render(http.StatusOK, "index", data)
}