package repository

import (
	"context"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"gorm.io/gorm"
)

type FileLoaderRepository interface {
	UploadFile(ctx context.Context, file *entity.File) error
}

type fileLoaderRepository struct {
	db *gorm.DB
}

func NewFileLoaderRepository(db *gorm.DB) FileLoaderRepository {
	return &fileLoaderRepository{db}
}

func (f *fileLoaderRepository) UploadFile(ctx context.Context, file *entity.File) error {
	return f.db.WithContext(ctx).Create(file).Error
}
