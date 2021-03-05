package entities

import (
	"github.com/DoNewsCode/core/otgorm"
	"gorm.io/gorm"
	"time"
)

func Migrations() []*otgorm.Migration {
	return []*otgorm.Migration{
		{
			ID: "202103042231",
			Migrate: func(db *gorm.DB) error {
				type User struct {
					gorm.Model
					Email         string    `gorm:"size:255;uniqueIndex;not null"`
					Username      string    `gorm:"size:255;uniqueIndex;not null"`
					Password      string    `gorm:"size:255"`
					EmailVerifyAt time.Time `gorm:"email_verify_at"`
					RememberToken string    `gorm:"remember_token"`
				}
				return db.AutoMigrate(&User{})
			},
			Rollback: func(db *gorm.DB) error {
				type User struct{}
				return db.Migrator().DropTable(&User{})
			},
		},
	}
}
