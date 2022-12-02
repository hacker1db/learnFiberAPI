package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type UpdateUser struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
