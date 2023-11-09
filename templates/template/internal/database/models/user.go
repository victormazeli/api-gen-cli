package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName  string    `json:"full_name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null""`
	Password  string    `json:"password" gorm:"type:varchar(255);not null""`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
