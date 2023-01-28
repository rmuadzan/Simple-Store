package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SearchProductHandler(ctx echo.Context) error {
	title := ctx.QueryParam("title")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 20
	
	products, totalRows, err := repositories.SearchProductByTitle(title, page, perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagination, _ := repositories.GetPaginationLinks(models.PaginationParams{
		Path:        "search",
		TotalRows:   totalRows,
		PerPage:     perPage,
		CurrentPage: page,
	})


	data := struct{
		Products *[]*models.Product
		Title string
		Length int
		Pagination models.PaginationLinks
	}{}

	data.Products = products
	data.Title = title
	data.Length = len(*products)
	data.Pagination = pagination

	return ctx.Render(http.StatusOK, "searchProducts", data)
}