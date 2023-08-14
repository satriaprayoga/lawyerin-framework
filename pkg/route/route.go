package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Route struct {
	R    *echo.Echo
	Port string
}

func New(port string, debug bool) Route {
	r := echo.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.CORS())
	if debug {
		r.Use(middleware.Logger())
	}
	r.Use(middleware.Recover())
	return Route{
		R:    r,
		Port: port,
	}
}
