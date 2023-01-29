package middlewares

import (
	"net/http"
	"simple-catalog-v2/repositories"

	"github.com/labstack/echo/v4"
)

func MiddlewareJWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/login" || ctx.Path() == "/signup" || ctx.Path() == "/auth" || ctx.Path() == "/logout" || ctx.Path() == "/forget-password" || ctx.Path() == "/forget-password/confirm" || ctx.Path() == "/forget-password/change" {
			return next(ctx)
		}

		token, ok := repositories.GetJwtTokenFromCookies(ctx)
		if !ok {
			return ctx.Redirect(http.StatusSeeOther, "/logout")
		}

		claims, err := repositories.GetUserClaims(token)
		if err != nil {
			return ctx.Redirect(http.StatusSeeOther, "/logout")
		}

		ctx.Set("userInfo", claims)
		return next(ctx)
	}
}