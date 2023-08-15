package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/interfaces/putusans"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
	"github.com/satriaprayoga/lawyerin-framework/usecases"
)

type PutusanController struct {
	putusanUC putusans.PutusanService
}

func NewPutusanController(s *data.Store) PutusanController {
	putusanService := usecases.NewPutusanUsecase(s)
	return PutusanController{putusanService}
}

func (a *PutusanController) Get(c echo.Context) error {
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
	putusanID, _ := strconv.Atoi(id)
	data, err := a.putusanUC.GetByID(ctx, putusanID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", data)
}

func (a *PutusanController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = web.Resp{R: c}
		//	logger  = logger.GetLogger()
		putusan = data.Putusan{}
	)

	httpCode, errMsg := web.BindAndValid(c, &putusan)
	if httpCode != 200 {
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}

	err := a.putusanUC.Create(ctx, &putusan)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", errMsg), nil)
	}

	return resp.Response(http.StatusOK, "ok", putusan)
}

func (a *PutusanController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp   = web.Resp{R: c}
		id     string
		update = data.Putusan{}
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
	putusanID, _ := strconv.Atoi(id)
	err := a.putusanUC.Update(ctx, putusanID, update)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", putusanID)
}

func (a *PutusanController) Delete(c echo.Context) error {
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
	putusanID, _ := strconv.Atoi(id)
	err := a.putusanUC.Delete(ctx, putusanID)
	if err != nil {
		return resp.ResponseError(http.StatusNotFound, "item not found", nil)
	}

	return resp.Response(http.StatusOK, "ok", nil)
}

func (a *PutusanController) TextSearch(c echo.Context) error {
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

	result, err := a.putusanUC.TextSearch(ctx, s.SearchTerm)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	return resp.ResponseList(http.StatusOK, "ok", result)
}
