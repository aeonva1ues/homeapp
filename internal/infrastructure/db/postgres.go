package db

import (
	"fmt"

	"github.com/aeonva1ues/homeapp/internal/entity"
	"github.com/aeonva1ues/homeapp/internal/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *infrastructure.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&entity.File{}, &entity.Category{}); err != nil {
		return nil, err
	}
	return db, nil
}
