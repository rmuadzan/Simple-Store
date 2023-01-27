package middlewares

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

func MiddlewareJWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/login" || ctx.Path() == "/signup" || ctx.Path() == "/auth" || ctx.Path() == "/logout" {
			return next(ctx)
		}

		token, ok := repositories.GetJwtTokenFromCookies(ctx)
		if !ok {
			return ctx.Redirect(http.StatusSeeOther, "/login")
		}

		claims, err := repositories.GetUserClaims(token)
		if err != nil {
			fmt.Println("gagal")
			return ctx.Redirect(http.StatusSeeOther, "/logout")
		}

		ctx.Set("userInfo", claims)
		return next(ctx)
	}
}