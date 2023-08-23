package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	firms "github.com/satriaprayoga/lawyerin-framework/interfaces/firm"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
	"github.com/satriaprayoga/lawyerin-framework/usecases"
)

type RadiusForm struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type FirmController struct {
	firmUC firms.FirmService
}

func NewFirmController(s *data.Store) FirmController {
	firmService := usecases.NewFirmUsecase(s)
	return FirmController{firmUC: firmService}
}

func (a *FirmController) Radius(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp = web.Resp{R: c}
		//logg  = logger.GetLogger()
		r      = RadiusForm{}
		result = web.ResponseModelList{}
	)

	httpCode, errMsg := web.BindAndValid(c, &r)
	if httpCode != 200 {
		//logg.Error("%v",errMsg)
		return resp.ResponseError(httpCode, fmt.Sprintf("%v", errMsg), nil)
	}
	result, err := a.firmUC.FindByRadius(ctx, r.Lat, r.Lng)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	return resp.ResponseList(http.StatusOK, "ok", result)

}
