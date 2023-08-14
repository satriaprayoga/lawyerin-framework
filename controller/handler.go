package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
)

var (
	e *echo.Echo
	s *data.Store
)

type Handlers struct {
	ArticleController ArticleController
}

func New(E *echo.Echo, S *data.Store) *Handlers {
	e = E
	s = S
	Static()
	return &Handlers{ArticleController: NewArticleController(s)}
}

func Static() {
	e.Static("/public", "public")
}

func (h *Handlers) Routes() {
	a := e.Group("/articles")
	a.GET("/", h.ArticleController.Get)
}
