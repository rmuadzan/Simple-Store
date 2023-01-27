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

	productsDataSlice := []*models.DisplayProductData{}
	for i := 0; i < len(*products); i++ {
		var product models.DisplayProductData
		user, _ := repositories.GetUserInfoByEmailOrId("", (*products)[i].UserID)
		product.Id = (*products)[i].Id
		product.Thumbnail = (*products)[i].Thumbnail
		product.Price = (*products)[i].Price
		product.FPrice = (*products)[i].FPrice
		product.DiscountPercentage = (*products)[i].DiscountPercentage
		product.Title = (*products)[i].Title
		product.StoreName = user.Fullname
		productsDataSlice = append(productsDataSlice, &product)
	}

	userInfo := repositories.GetUserClaimsFromContext(ctx)
	data := struct{
		Products *[]*models.DisplayProductData
		Length int
		UserStatus string
		Pagination models.PaginationLinks
	}{}

	data.Products = &productsDataSlice
	data.Length = len(*products)
	data.UserStatus = userInfo.Status
	data.Pagination = pagination

	return ctx.Render(http.StatusOK, "allProducts", data)
}