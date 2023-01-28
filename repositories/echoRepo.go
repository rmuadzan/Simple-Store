package repositories

import (
	"fmt"
	"net/http"
	"os"
	"simple-catalog-v2/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
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

// ECHO SECURE COOKIE
var _ = godotenv.Load()
var hashKey = os.Getenv("COOKIE_HASH_KEY")
var blockKey = os.Getenv("COOKIE_BLOCK_KEY")
var sc = securecookie.New([]byte(hashKey), []byte(blockKey))

func SetCookie(ctx echo.Context, name string, data interface{}) error {
	encodedValue, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	c := new(http.Cookie)
	c.Name = "jwt"
	c.Value = encodedValue
	c.Path = "/"
	c.Secure = false
	c.HttpOnly = true

	http.SetCookie(ctx.Response(), c)
	return nil
}

func GetCookie(ctx echo.Context, name string) (string, error) {
	c, err := ctx.Cookie(name)
	if err != nil {
		return "", err
	}

	var decodedValue string
	err = sc.Decode(name, c.Value, &decodedValue)
	return decodedValue, err
}

func DeleteCookie(ctx echo.Context, name string) {
	c := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(ctx.Response(), c)
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