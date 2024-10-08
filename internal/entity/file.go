package entity

import (
	"time"
)

type File struct {
	ID 		   uint      `gorm:"primary_key";"AUTO_INCREMENT" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Path       string    `gorm:"not null" json:"path"`
	SizeMB     float64     `gorm:"not null" json:"size_mb"`
	UploadedAt time.Time `gorm:"not null" json:"uploaded_at"`
	Removed    bool      `gorm:"default:false" json:"removed"`
}
