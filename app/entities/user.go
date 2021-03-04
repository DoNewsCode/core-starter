package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint           `json:"id",gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at",gorm:"index"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	Password      string         `json:"-"`
	EmailVerifyAt time.Time      `json:"email_verify_at"`
	RememberToken string         `json:"remember_token"`
}
