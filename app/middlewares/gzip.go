package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Gzip() echo.MiddlewareFunc {
	return middleware.Gzip()
}