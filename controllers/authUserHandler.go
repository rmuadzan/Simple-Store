package controllers

import (
	"net/http"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// POST "/auth"
func AuthUserHandler(ctx echo.Context) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	userInfo, err := repositories.GetUserInfoByEmailOrId(email, 0)
	if err != nil {
		return ctx.Redirect(http.StatusSeeOther, "/login")
	} 
	actualPassword := userInfo.Password

	err = bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(password))
	if err != nil {
		return ctx.Redirect(http.StatusSeeOther, "/login")
	}

	claims, err := repositories.GenerateUserClaims(userInfo.Id, userInfo.Fullname, userInfo.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := repositories.GenerateUserToken(claims)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := repositories.SetCookie(ctx, "jwt", token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/")
}