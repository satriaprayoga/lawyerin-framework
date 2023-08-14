package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/articles"
	"github.com/satriaprayoga/lawyerin-framework/usecases"
)

type ArticleController struct {
	articleUC articles.ArticleService
}

func NewArticleController(s *data.Store) ArticleController {
	articleService := usecases.NewArticleUsecase(s)
	return ArticleController{articleService}
}

func (a *ArticleController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Article OK")
}
