package fileloader

import (
	"mime/multipart"
	"sync"
	"time"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/gin-gonic/gin"
)

func LoadFile(ctx *gin.Context, file *multipart.FileHeader, dst string, c chan *entity.File, w *sync.WaitGroup) {
	ctx.SaveUploadedFile(file, dst)
	filename, filepath := file.Filename, dst+file.Filename
	c <- &entity.File{
		Name:       filename,
		CategoryID: 1,
		Path:       filepath,
		SizeMB:     float64(file.Size) / 1024,
		UploadedAt: time.Now(),
	}
	w.Done()
}
