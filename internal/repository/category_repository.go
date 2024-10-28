package repository

import (
	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx *gin.Context, category *entity.Category) error
	SelectCategories(ctx *gin.Context) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) CreateCategory(ctx *gin.Context, category *entity.Category) error {
	return c.db.WithContext(ctx).Create(category).Error
}

func (c *categoryRepository) SelectCategories(ctx *gin.Context) error {
	var categories []entity.Category
	return c.db.WithContext(ctx).Find(&categories).Error
}
