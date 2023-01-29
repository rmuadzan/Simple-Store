package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AllProductsHandler(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 20
	
	products, totalRows, err := repositories.GetAllProducts(page, perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagination, _ := repositories.GetPaginationLinks(models.PaginationParams{
		Path:        "products",
		TotalRows:   totalRows,
		PerPage:     perPage,
		CurrentPage: page,
	})

	userInfo := repositories.GetUserClaimsFromContext(ctx)
	data := struct{
		Products *[]*models.Product
		Length int
		UserStatus string
		Pagination models.PaginationLinks
	}{}

	data.Products = products
	data.Length = len(*products)
	data.UserStatus = userInfo.Status
	data.Pagination = pagination

	return ctx.Render(http.StatusOK, "allProducts", data)
}