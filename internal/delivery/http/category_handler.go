package http

import (
	"net/http"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
	log             *logrus.Logger
}

func NewCategoryHandler(router *gin.Engine, usecase usecase.CategoryUseCase, log *logrus.Logger) {
	handler := &CategoryHandler{
		categoryUseCase: usecase,
		log:             log,
	}
	categoryRoutes := router.Group("/category")
	{
		categoryRoutes.GET("/", handler.SelectAll)
		categoryRoutes.POST("/", handler.Create)
	}
}

func (c *CategoryHandler) SelectAll(ctx *gin.Context) {
	c.categoryUseCase.SelectCategories(ctx)
	ctx.Status(http.StatusOK)
}

func (c *CategoryHandler) Create(ctx *gin.Context) {
	var category entity.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		c.log.Errorf("error during parsing json while creation category: %s", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err := c.categoryUseCase.CreateCategory(ctx, &category); err != nil {
		c.log.Errorf("error during creation category usecase: %s", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusCreated)
}
