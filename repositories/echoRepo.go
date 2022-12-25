package repositories

import (
	"fmt"
	"net/http"
	"simple-catalog-v2/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func NewRenderer(location string, debug bool) *models.Renderer {
	tpl := new(models.Renderer)
	tpl.Location = location
	tpl.Debug = debug

	tpl.ReloadTemplate()

	return tpl
}

func ErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		code = report.Code
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			}
			break
		}
	}
	   

	data := struct{
		Message string
	}{
		report.Message.(string),
	}
	ctx.Render(code, "error", data)
}

func MakeLogEntry(ctx echo.Context) *log.Entry {
	if ctx == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2023-01-01 00:00:001"),
		})
	}

	return log.WithFields(log.Fields{
		"at": time.Now().Format("2023-01-01 00:00:001"),
		"method": ctx.Request().Method,
		"uri": ctx.Request().URL.String(),
		"ip": ctx.Request().RemoteAddr,
	})
}