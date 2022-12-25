package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

func SearchProductHandler(ctx echo.Context) error {
	title := ctx.QueryParam("title")
	products, err := repositories.SearchProductByTitle(title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	data := struct{
		Products *[]*models.Product
		Title string
		Length int
	}{}

	data.Products = products
	data.Title = title
	data.Length = len(*products)

	return ctx.Render(http.StatusOK, "searchProducts", data)
}