package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductHandler(ctx echo.Context) error {
	data := new(models.Product)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.Init()
	data.FPrice = float64(int(data.Price * (1 - data.DiscountPercentage / 100.0) * 100)) / 100
	data.UserID = repositories.GetUserClaimsFromContext(ctx).Id

	if err := repositories.CreateProduct(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserProductsHandler(ctx)
	return nil
} 

func GetProductByIdHandler(ctx echo.Context) error {
	data := struct{
		IsOwner bool
		Product models.Product
		UserStatus string
	}{}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	product, err := repositories.GetProductById(id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusNotFound, notFoundMessage)
	}
	if userInfo.Id == product.UserID  || userInfo.Status == "admin" {
		data.IsOwner = true
	}
	data.Product = product
	data.UserStatus = userInfo.Status
		
	return ctx.Render(http.StatusOK, "detailProduct", data)
}

func UpdateProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusNotFound, notFoundMessage)
	}

	data := new(models.Product)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.FPrice = float64(int(data.Price * (1 - data.DiscountPercentage / 100.0) * 100)) / 100

	if err := repositories.UpdateProductById(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	GetProductByIdHandler(ctx)
	return nil
}

func DeleteProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusNotFound, notFoundMessage)
	}
	
	err = repositories.DeleteProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/products")
}