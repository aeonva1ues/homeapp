package main

import (
	"github.com/aeonva1ues/homeapp/internal/delivery/http"
	"github.com/aeonva1ues/homeapp/internal/infrastructure"
	"github.com/aeonva1ues/homeapp/internal/infrastructure/db"
	"github.com/aeonva1ues/homeapp/internal/infrastructure/logger"
	"github.com/aeonva1ues/homeapp/internal/repository"
	"github.com/aeonva1ues/homeapp/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()

	cfg, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	fileLoaderRepository := repository.NewFileLoaderRepository(db)

	fileLoaderUseCase := usecase.NewFileLoaderUseCase(fileLoaderRepository, cfg.Uploads.Path)

	router := gin.Default()
	router.LoadHTMLGlob("html/*")

	http.NewFileLoaderHandler(router, fileLoaderUseCase, log)

	if err := router.Run(cfg.Server.IP + cfg.Server.Port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
