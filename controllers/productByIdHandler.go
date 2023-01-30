package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
) 

// GET "/products/:id"
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

// POST "/products/:id"
func UpdateProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	product, err := repositories.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	if product.UserID != userInfo.Id || userInfo.Status != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Sorry you don't have access to access this product")
	}

	data := new(models.Product)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.FPrice = float64(int(data.Price * (1 - data.DiscountPercentage / 100.0) * 100)) / 100

	if err := repositories.UpdateProduct(data, userInfo.Id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	GetProductByIdHandler(ctx)
	return nil
}

// GET "/products/:id/edit"
func EditProductHandler(ctx echo.Context) error {
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

	if product.UserID != userInfo.Id && userInfo.Status != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Sorry you don't have authority to access this product")
	}

	return ctx.Render(http.StatusOK, "editProduct", product)
}

// POST "products/:id/delete"
func DeleteProductByIdHandler(ctx echo.Context) error {
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

	if product.UserID != userInfo.Id && userInfo.Status != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Sorry you don't have authority to access this product")
	}
	
	err = repositories.DeleteProduct(id, userInfo.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/products")
}

// GET "/products/:id/order"
func AddOrderProduct(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	product, err := repositories.GetProductById(id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No product match with id %d", id)
		return echo.NewHTTPError(http.StatusNotFound, notFoundMessage)
	}

	return ctx.Render(http.StatusOK, "orderProduct", product)
}