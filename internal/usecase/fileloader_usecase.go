package usecase

import (
	"context"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/repository"
)

type FileLoaderUseCase interface {
	UploadFile(ctx context.Context, file *entity.File) error
}

type fileLoaderUseCase struct {
	repository repository.FileLoaderRepository
}

func NewFileLoaderUseCase(fileloaderRepository repository.FileLoaderRepository) FileLoaderUseCase {
	return &fileLoaderUseCase{repository: fileloaderRepository}
}

func (f *fileLoaderUseCase) UploadFile(ctx context.Context, file *entity.File) error {
	return f.repository.UploadFile(ctx, file)
}
