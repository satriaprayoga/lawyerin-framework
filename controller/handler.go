package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/pkg/filesystem"
)

var (
	e  *echo.Echo
	s  *data.Store
	fs *filesystem.FS
)

type Handlers struct {
	ArticleController   ArticleController
	PutusanController   PutusanController
	PeraturanController PeraturanController
	FileController      FileController
	FirmController      FirmController
}

func New(E *echo.Echo, S *data.Store, FS *filesystem.FS) *Handlers {
	e = E
	s = S
	fs = FS
	Static()
	return &Handlers{ArticleController: NewArticleController(s),
		PutusanController:   NewPutusanController(s),
		PeraturanController: NewPeraturanController(s),
		FileController:      NewFileController(*fs),
		FirmController:      NewFirmController(s)}
}

func Static() {
	e.Static("/public", "public")
}

func (h *Handlers) Routes() {
	a := e.Group("/article")
	a.GET("/:id", h.ArticleController.Get)
	a.PUT("/:id", h.ArticleController.Update)
	a.DELETE("/:id", h.ArticleController.Delete)
	a.POST("/create", h.ArticleController.Create)
	a.GET("/search", h.ArticleController.TextSearch)

	b := e.Group("/putusan")
	b.GET("/:id", h.PutusanController.Get)
	b.PUT("/:id", h.PutusanController.Update)
	b.DELETE("/:id", h.PutusanController.Delete)
	b.POST("/create", h.PutusanController.Create)
	b.GET("/search", h.PutusanController.TextSearch)

	c := e.Group("/peraturan")
	c.GET("/:id", h.PeraturanController.Get)
	c.PUT("/:id", h.PeraturanController.Update)
	c.DELETE("/:id", h.PeraturanController.Delete)
	c.POST("/create", h.PeraturanController.Create)
	c.GET("/search", h.PeraturanController.TextSearch)

	d := e.Group("/files")
	d.PUT("/upload", h.FileController.Upload)

	e := e.Group("/firms")
	e.GET("/radius", h.FirmController.Radius)
}
