package models

import (
	"context"
	"html/template"
	"io"

	"github.com/go-playground/validator/v10"
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

