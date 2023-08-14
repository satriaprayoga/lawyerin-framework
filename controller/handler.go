package controller

import "github.com/labstack/echo/v4"

var e *echo.Echo

type Handler struct {
}

func New(E *echo.Echo) *Handler {
	e = E
	Static()
	return &Handler{}
}

func Static() {
	e.Static("/public", "public")
}
