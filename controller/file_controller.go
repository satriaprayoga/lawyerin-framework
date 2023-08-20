package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/pkg/filesystem"
	"github.com/satriaprayoga/lawyerin-framework/pkg/web"
)

type FileController struct {
	fs filesystem.FS
}

func NewFileController(fs filesystem.FS) FileController {
	return FileController{fs: fs}
}

func (f *FileController) Upload(c echo.Context) error {

	var (
		resp = web.Resp{R: c}
	)
	filename, err := f.getUploadFile(c)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), err)
	}

	// err = f.fs.Put("./tmp/"+filename, "test")
	// if err != nil {
	// 	return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), err)
	// }

	return resp.Response(http.StatusOK, "OK", filename)

}

func (f *FileController) getUploadFile(c echo.Context) (string, error) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	defer src.Close()

	// Destination
	errDir := os.MkdirAll("./public/profile/user/", 0777)
	if err != nil {
		return "", errDir
	}
	dst, err := os.Create("./public/profile/user/" + file.Filename)
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Printf("%v", err)
		return "", err
	}

	return file.Filename, err

}
