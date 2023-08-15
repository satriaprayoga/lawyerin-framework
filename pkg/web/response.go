package web

import "github.com/labstack/echo/v4"

type Resp struct {
	R echo.Context
}

type ResponseModel struct {
	Message string      `jsong:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelList struct {
	Page         int         `json:"page"`
	Total        int         `json:"total"`
	LastPage     int         `json:"last_page"`
	DefineSize   string      `json:"define_size"`
	DefineColumn string      `json:"define_column"`
	AllColumn    string      `json:"all_column"`
	Data         interface{} `json:"data"`
	Msg          string      `json:"message"`
}

func (e Resp) Response(httpCode int, errMsg string, data interface{}) error {

	response := ResponseModel{
		Message: errMsg,
		Data:    data,
	}

	return e.R.JSON(httpCode, response)
}

func (e Resp) ResponseError(httpCode int, errMsg string, data interface{}) error {

	response := ResponseModel{

		Message: errMsg,
		Data:    data,
	}

	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseErrorList :
func (e Resp) ResponseErrorList(httpCode int, errMsg string, data ResponseModelList) error {

	data.Msg = errMsg

	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}

// ResponseList :
func (e Resp) ResponseList(httpCode int, errMsg string, data ResponseModelList) error {

	data.Msg = errMsg

	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}
