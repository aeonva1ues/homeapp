package usecase

import (
	"mime/multipart"
	"sync"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/fileloader"
	"github.com/aeonva1ues/homeapp/internal/repository"
	"github.com/gin-gonic/gin"
)

type FileLoaderUseCase interface {
	UploadFiles(ctx *gin.Context, files *[]*multipart.FileHeader) error
}

type fileLoaderUseCase struct {
	repository repository.FileLoaderRepository
	dst        string
}

func NewFileLoaderUseCase(fileloaderRepository repository.FileLoaderRepository, dst string) FileLoaderUseCase {
	return &fileLoaderUseCase{repository: fileloaderRepository, dst: dst}
}

func (f *fileLoaderUseCase) UploadFiles(ctx *gin.Context, files *[]*multipart.FileHeader) error {
	var w sync.WaitGroup
	uploadedFiles := make(chan *entity.File)
	for _, file := range *files {
		w.Add(1)
		go fileloader.LoadFile(ctx, file, f.dst+file.Filename, uploadedFiles, &w)
	}
	go func() {
		w.Wait()
		close(uploadedFiles)
	}()
	var filesToInsert []entity.File
	for uploadedFile := range uploadedFiles {
		filesToInsert = append(filesToInsert, *uploadedFile)
	}
	return f.repository.UploadFiles(ctx, &filesToInsert)
}
