package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProductHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "addProduct", nil)
}