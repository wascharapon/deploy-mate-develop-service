package api

import (
	"github.com/labstack/echo/v4"
)

func BindQueryParams(c echo.Context, v any) error {
	return (&echo.DefaultBinder{}).BindQueryParams(c, v)
}
