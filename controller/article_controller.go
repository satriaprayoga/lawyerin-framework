package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/articles"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
	"github.com/satriaprayoga/lawyerin-framework/usecases"
)

type SearchForm struct {
	SearchTerm string `json:"search_term"`
}

type ArticleController struct {
	articleUC articles.ArticleService
}

func NewArticleController(s *data.Store) ArticleController {
	articleService := usecases.NewArticleUsecase(s)
	return ArticleController{articleService}
}

func (a *ArticleController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp = web.Resp{R: c}
		id   string
	)

	id = c.Param("id")
	if id == "" {
		return resp.ResponseError(http.StatusBadRequest, "invalid or empty path", nil)
	}
	articleID, _ := strconv.Atoi(id)
	data, err := a.articleUC.GetByID(ctx, articleID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", data)
}

func (a *ArticleController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = web.Resp{R: c}
		//	logger  = logger.GetLogger()
		article = data.Article{}
	)

	httpCode, errMsg := web.BindAndValid(c, &article)
	if httpCode != 200 {
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}

	err := a.articleUC.Create(ctx, &article)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", errMsg), nil)
	}

	return resp.Response(http.StatusOK, "ok", article)
}

func (a *ArticleController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp   = web.Resp{R: c}
		id     string
		update = data.Article{}
	)

	id = c.Param("id")
	if id == "" {
		return resp.ResponseError(http.StatusBadRequest, "invalid or empty path", nil)
	}
	httpCode, errMsg := web.BindAndValid(c, &update)
	if httpCode != 200 {
		//logg.Error("%v",errMsg)
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}
	articleID, _ := strconv.Atoi(id)
	err := a.articleUC.Update(ctx, articleID, update)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", articleID)
}

func (a *ArticleController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp = web.Resp{R: c}
		id   string
	)

	id = c.Param("id")
	if id == "" {
		return resp.ResponseError(http.StatusBadRequest, "invalid or empty path", nil)
	}
	articleID, _ := strconv.Atoi(id)
	err := a.articleUC.Delete(ctx, articleID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", nil)
}

func (a *ArticleController) TextSearch(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = web.Resp{R: c}
		//logg  = logger.GetLogger()
		s      = SearchForm{}
		result = web.ResponseModelList{}
	)

	httpCode, errMsg := web.BindAndValid(c, &s)
	if httpCode != 200 {
		//logg.Error("%v",errMsg)
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}

	result, err := a.articleUC.TextSearch(ctx, s.SearchTerm)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	return resp.ResponseList(http.StatusOK, "ok", result)
}
