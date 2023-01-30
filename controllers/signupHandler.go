package controllers

import (
	"net/http"
	"simple-catalog-v2/models"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// GET "/signup"
func SignUpPage(ctx echo.Context) error {
	if _, err := ctx.Cookie("jwt"); err == nil {
		return ctx.Redirect(http.StatusSeeOther, "/")
	}
	
	return ctx.Render(http.StatusOK, "signup", nil)
}

// POST "/signup"
func SignUpHandler(ctx echo.Context) error {
	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := ctx.Validate(&user); err != nil {
		return err
	}

	if user.Status != "user" && user.Status != "store" {
		return ctx.Redirect(http.StatusSeeOther, "/signup")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(ctx.FormValue("password")), 8)
	if err != nil { 
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password=string(password)
	err = repositories.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	AuthUserHandler(ctx)
	return nil
}

