package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserProductsHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 20

	products, totalRows, err := repositories.GetUserProducts(userInfo.Id, page, perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagination, _ := repositories.GetPaginationLinks(models.PaginationParams{
		Path:        "my-products",
		TotalRows:   totalRows,
		PerPage:     perPage,
		CurrentPage: page,
	})

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

	return ctx.Render(http.StatusOK, "userProducts", data)
}