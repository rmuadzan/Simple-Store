package models

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// ECHO RENDERER
type Renderer struct {
	template *template.Template
	Location string
	Debug bool
}

func (t *Renderer) ReloadTemplate() {
	t.template = template.Must(template.ParseGlob(t.Location))
}

func (t *Renderer) Render (
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.Debug {
		t.ReloadTemplate()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

// ECHO CUSTOM CONTEXT
type ContextValue struct {
    echo.Context
}

func (ctx *ContextValue) Get(key string) interface{} {
    val := ctx.Context.Get(key)
    if val != nil {
        return val
    }
    return ctx.Request().Context().Value(key)
}

func (ctx *ContextValue) Set(key string, val interface{}) {
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), key, val)))
}

// ECHO VALIDATOR
type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
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

	ctx.SetCookie(c)
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

	ctx.SetCookie(c)
}
   