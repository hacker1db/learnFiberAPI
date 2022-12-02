package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID           uint `json:"id" gorm:"PrimaryKey"`
	CreatedAt    time.Time
	Name         string `json:"Name"`
	SerialNumber uuid.UUID
}

type RoutesProduct struct {
	ID           uint `json:"ID"`
	SerialNumber uuid.UUID
	Name         string `json:"Name"`
}

type UpdateProduct struct {
	Name string `json:"Name"`
}
