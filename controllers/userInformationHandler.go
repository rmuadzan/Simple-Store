package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

// GET "/profile"
func GetUserInformationHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)

	user, err := repositories.GetUserInfoByEmailOrId("", userInfo.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var data models.DisplayUserData
	data.Avatar = user.Avatar
	data.Fullname = user.Fullname
	data.Username = user.Username
	data.Email = user.Email
	data.Gender = user.Gender
	data.Status = user.Status
		
	return ctx.Render(http.StatusOK, "userInformation", data)
}

// POST "/profile"
func UpdateUserInformationHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)
	var user models.DisplayUserData

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	if user.Status != "user" && user.Status != "store" {
		return ctx.Redirect(http.StatusSeeOther, "/profile/edit")
	}

	err := repositories.UpdateUser(userInfo.Id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims, err := repositories.GenerateUserClaims(userInfo.Id, user.Fullname, user.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	token, err := repositories.GenerateUserToken(claims)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	repositories.SetCookie(ctx, "jwt", token)

	GetUserInformationHandler(ctx)
	return nil
}

// GET "/profile/edit"
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