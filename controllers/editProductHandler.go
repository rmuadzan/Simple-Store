package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func EditProductHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	product, err := repositories.GetProductById(id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusInternalServerError, notFoundMessage)
	}

	return ctx.Render(http.StatusOK, "editProduct", product)
}