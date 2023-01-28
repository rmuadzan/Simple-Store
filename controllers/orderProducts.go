package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddOrderProduct(ctx echo.Context) error {
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

// Get /my-order
func UserOrderHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)
	orders, err := repositories.GetUserOrders(userInfo.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}


	data := struct{
		Orders *[]*models.Order
		Length int
		UserStatus string
	}{}

	data.Orders = orders
	data.Length = len(*orders)
	data.UserStatus = userInfo.Status

	return ctx.Render(http.StatusOK, "userOrder", data)
}

// Post /my-order
func OrderProductHandler(ctx echo.Context) error {
	data := new(models.Order)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.UserID = repositories.GetUserClaimsFromContext(ctx).Id

	if err := repositories.CreateOrder(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserOrderHandler(ctx)
	return nil
} 

func GetOrderByIdHandler(ctx echo.Context) (error) {
	var order models.Order
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	order, err = repositories.GetOrderById(id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No order match with id %d", id)
		return echo.NewHTTPError(http.StatusInternalServerError, notFoundMessage)
	}
		
	return ctx.Render(http.StatusOK, "orderDetail", order)
}

func DeleteOrderByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		notFoundMessage := fmt.Sprintf("No order match with id %d", id)
		return echo.NewHTTPError(http.StatusBadRequest, notFoundMessage)
	}
	
	err = repositories.DeleteOrderById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/my-order")
}