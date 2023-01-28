package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

func EditUserInformationHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)
	user, err := repositories.GetUserInfoByEmailOrId("", userInfo.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var data models.DisplayUserData
	data.Fullname = user.Fullname
	data.Username = user.Username
	data.Gender = user.Gender
	data.Status = user.Status

	return ctx.Render(http.StatusOK, "editProfile", data)
}