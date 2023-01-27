package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func OrderProduct(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	product, err := repositories.GetProductById(id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusInternalServerError, notFoundMessage)
	}

	return ctx.Render(http.StatusOK, "orderProduct", product)
}

func OrderProductHandler(ctx echo.Context) error {
	data := new(models.Product)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.Init()
	data.FPrice = float64(int(data.Price * (1 - data.DiscountPercentage / 100.0) * 100)) / 100
	data.UserID = 1

	if err := repositories.CreateProduct(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserProductsHandler(ctx)
	return nil
} 