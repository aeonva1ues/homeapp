package http

import (
	"net/http"

	"github.com/aeonva1ues/homeapp/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FileLoaderHandler struct {
	fileLoaderUseCase usecase.FileLoaderUseCase
	log               *logrus.Logger
}

func NewFileLoaderHandler(router *gin.Engine, usecase usecase.FileLoaderUseCase, log *logrus.Logger) {
	handler := &FileLoaderHandler{
		fileLoaderUseCase: usecase,
		log:               log,
	}
	fileLoaderRoutes := router.Group("/file")
	{
		fileLoaderRoutes.GET("/", handler.RenderMain)
		fileLoaderRoutes.POST("/uploads", handler.UploadFiles)
	}
}

func (h *FileLoaderHandler) RenderMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h *FileLoaderHandler) UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	if err := h.fileLoaderUseCase.UploadFiles(c, &files); err != nil {
		h.log.Errorf("error during loading files x%d - %s", len(files), err)
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Status(http.StatusCreated)
}
