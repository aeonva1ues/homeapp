package repository

import (
	"context"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"gorm.io/gorm"
)

type FileLoaderRepository interface {
	UploadFiles(ctx context.Context, file *[]entity.File) error
}

type fileLoaderRepository struct {
	db *gorm.DB
}

func NewFileLoaderRepository(db *gorm.DB) FileLoaderRepository {
	return &fileLoaderRepository{db}
}

func (f *fileLoaderRepository) UploadFiles(ctx context.Context, files *[]entity.File) error {
	if len(*files) == 0 {
		return ErrNoContent
	}
	return f.db.WithContext(ctx).Create(files).Error
}
