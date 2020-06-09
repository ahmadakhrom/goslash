package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func Logger () echo.MiddlewareFunc {
	out, err := os.Create("logs/logs.txt")
	if err != nil {
		out = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  "method=${method}, uri=${uri}, Latency:${latency_human}, status=${status}\n",
		Output:  out,
	})

}
