package http

import (
	"net/http"
	"sync"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/fileloader"
	"github.com/aeonva1ues/homeapp/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FileLoaderHandler struct {
	fileLoaderUseCase usecase.FileLoaderUseCase
	log               *logrus.Logger
	dst				  string
}

func NewFileLoaderHandler(router *gin.Engine, usecase usecase.FileLoaderUseCase, log *logrus.Logger, dst string) {
	handler := &FileLoaderHandler{
		fileLoaderUseCase: usecase,
		log:               log,
		dst:			   dst,
	}
	fileLoaderRoutes := router.Group("/file")
	{
		fileLoaderRoutes.GET("/", handler.RenderMain)
		fileLoaderRoutes.POST("/uploads", handler.UploadFile)
	}
}

func (h *FileLoaderHandler) RenderMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h *FileLoaderHandler) UploadFile(c *gin.Context) {
	var w sync.WaitGroup
	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) == 0 {
		c.Status(http.StatusNoContent)
		return
	}
	uploadedFiles := make(chan *entity.File)
	for _, file := range files {
		h.log.Info("got file " + file.Filename)
		w.Add(1)
		go fileloader.LoadFile(c, file, h.dst + file.Filename, uploadedFiles, &w)
	}
	go func() {
		w.Wait()
		close(uploadedFiles)
	}()
	for uploadedFile := range uploadedFiles {
		h.fileLoaderUseCase.UploadFile(c.Request.Context(), uploadedFile)
		h.log.Info("finish work with " + uploadedFile.Name)
	}
	c.Status(http.StatusOK)
}
