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

	cfg, err := infrastructure.LoadEnvConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	db, err := db.NewPostgresDB(cfg.GetDatabase())
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	fileLoaderRepository := repository.NewFileLoaderRepository(db)

	fileLoaderUseCase := usecase.NewFileLoaderUseCase(fileLoaderRepository, cfg.GetUploadsDir())

	router := gin.Default()
	router.LoadHTMLGlob("html/*")

	http.NewFileLoaderHandler(router, fileLoaderUseCase, log)

	if err := router.Run(cfg.GetServerHost()); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
