package controllers

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get "/my-orders"
func UserOrderHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 20

	orders, totalRows, err := repositories.GetUserOrders(userInfo.Id, page, perPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagination, _ := repositories.GetPaginationLinks(models.PaginationParams{
		Path:        "my-orders",
		TotalRows:   totalRows,
		PerPage:     perPage,
		CurrentPage: page,
	})

	data := struct{
		Orders *[]*models.Order
		Length int
		UserStatus string
		Pagination models.PaginationLinks
	}{}

	data.Orders = orders
	data.Length = len(*orders)
	data.UserStatus = userInfo.Status
	data.Pagination = pagination

	return ctx.Render(http.StatusOK, "userOrder", data)
}

// Post "/my-orders"
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

// GET "my-orders/:id"
func GetOrderByIdHandler(ctx echo.Context) (error) {
	var order models.Order
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userInfo := repositories.GetUserClaimsFromContext(ctx)
	order, err = repositories.GetOrder(id, userInfo.Id)
	if err != nil {
		notFoundMessage := fmt.Sprintf("No order match with id %d", id)
		return echo.NewHTTPError(http.StatusNotFound, notFoundMessage)
	}
		
	return ctx.Render(http.StatusOK, "orderDetail", order)
}

// POST "my-orders/:id/delete"
func DeleteOrderByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	err = repositories.DeleteOrder(id, userInfo.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/my-orders")
}