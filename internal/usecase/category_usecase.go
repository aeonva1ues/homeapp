package usecase

import (
	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/repository"
	"github.com/gin-gonic/gin"
)

type CategoryUseCase interface {
	CreateCategory(ctx *gin.Context, category *entity.Category) error
	SelectCategories(ctx *gin.Context) error
}

type categoryUseCase struct {
	repository repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repository: categoryRepository}
}

func (c *categoryUseCase) CreateCategory(ctx *gin.Context, category *entity.Category) error {
	return c.repository.CreateCategory(ctx, category)
}

func (c *categoryUseCase) SelectCategories(ctx *gin.Context) error {
	return c.repository.SelectCategories(ctx)
}
