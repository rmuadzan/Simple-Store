package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

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

func UpdateUserInformationHandler(ctx echo.Context) error {
	userInfo := repositories.GetUserClaimsFromContext(ctx)
	var user models.DisplayUserData

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// user.Fullname=ctx.FormValue("fullname")
	// user.Username=ctx.FormValue("username")
	// user.Email=ctx.FormValue("email")
	// user.Gender=ctx.FormValue("gender")
	// user.Status=ctx.FormValue("status")
	
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

	c := &http.Cookie{}
	c.Name = "jwt"
	c.Value = token
	ctx.SetCookie(c)

	GetUserInformationHandler(ctx)
	return nil
}