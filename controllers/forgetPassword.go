package controllers

import (
	"net/http"
	"simple-catalog-v2/repositories"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ForgetPasswordHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "forgetPassword", nil)
}

func ForgetPasswordTokenHandler(ctx echo.Context) error {
	data := struct{
		Email string
	}{ctx.FormValue("email")}
	return ctx.Render(http.StatusOK, "forgetPasswordToken", data)
}

func CreateRefreshTokenHandler(ctx echo.Context) error {
	email := ctx.FormValue("email")
	userInfo, err := repositories.GetUserInfoByEmailOrId(email, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	if userInfo.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	token, err := bcrypt.GenerateFromPassword([]byte(email + time.Now().String()), 8)
	if err != nil { 
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = repositories.SetUserRefreshToken(email, string(token))
	if err != nil { 
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ForgetPasswordTokenHandler(ctx)
	return nil
}

func ChangePasswordHandler(ctx echo.Context) error {
	data := struct{
		Email string
	}{ctx.FormValue("email")}
	return ctx.Render(http.StatusOK, "changePassword", data)
}

func ValidateRefreshToken(ctx echo.Context) error {
	inputToken := ctx.FormValue("token")
	email := ctx.FormValue("email")
	actualToken, err := repositories.GetUserRefreshToken(email)

	if err != nil { 
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if inputToken != actualToken {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	ChangePasswordHandler(ctx)
	return nil
}

func UpdateUserPassword(ctx echo.Context) error {
	password := ctx.FormValue("password")
	email := ctx.FormValue("email")
	err := repositories.SetUserPassword(email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, "/login")
}