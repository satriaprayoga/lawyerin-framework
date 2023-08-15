package web

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/satriaprayoga/lawyerin-framework/pkg/logger"

	"github.com/labstack/echo/v4"
)

func BindAndValid(c echo.Context, form interface{}) (int, string) {
	var logg = logger.GetLogger()
	err := c.Bind(form)
	if err != nil {
		logg.Error("%v", err)
		return http.StatusBadRequest, "invalid request parameter"
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, "internal server error"
	}
	if !check {
		return http.StatusBadRequest, MarkErrors(valid.Errors)
	}
	return http.StatusOK, "ok"
}

func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res = fmt.Sprintf("%s %s", err.Key, err.Message)
	}
	return res
}
