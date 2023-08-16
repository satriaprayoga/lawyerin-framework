package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/peraturans"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
	"github.com/satriaprayoga/lawyerin-framework/usecases"
)

type PeraturanController struct {
	peraturanUC peraturans.PeraturanService
}

func NewPeraturanController(s *data.Store) PeraturanController {
	peraturanService := usecases.NewPeraturanUsecase(s)
	return PeraturanController{peraturanService}
}

func (a *PeraturanController) Get(c echo.Context) error {
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
	peraturanID, _ := strconv.Atoi(id)
	data, err := a.peraturanUC.GetByID(ctx, peraturanID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", data)
}

func (a *PeraturanController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = web.Resp{R: c}
		//	logger  = logger.GetLogger()
		peraturan = data.Peraturan{}
	)

	httpCode, errMsg := web.BindAndValid(c, &peraturan)
	if httpCode != 200 {
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}

	err := a.peraturanUC.Create(ctx, &peraturan)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", errMsg), nil)
	}

	return resp.Response(http.StatusOK, "ok", peraturan)
}

func (a *PeraturanController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp   = web.Resp{R: c}
		id     string
		update = data.Peraturan{}
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
	peraturanID, _ := strconv.Atoi(id)
	update.Slug = update.Category + " " + update.Bidang + " " + update.SubBidang + " " + update.Creator
	err := a.peraturanUC.Update(ctx, peraturanID, update)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", peraturanID)
}

func (a *PeraturanController) Delete(c echo.Context) error {
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
	peraturanID, _ := strconv.Atoi(id)
	err := a.peraturanUC.Delete(ctx, peraturanID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", nil)
}

func (a *PeraturanController) TextSearch(c echo.Context) error {
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

	result, err := a.peraturanUC.TextSearch(ctx, s.SearchTerm)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	return resp.ResponseList(http.StatusOK, "ok", result)
}
