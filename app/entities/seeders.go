package entities

import (
	"fmt"
	"github.com/DoNewsCode/core/otgorm"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Seeders() []*otgorm.Seed {
	return []*otgorm.Seed{
		{
			Name: "seeding users",
			Run: func(db *gorm.DB) error {
				for i := 0; i < 10; i++ {
					u := RandStringRunes(i + 3)
					db.Create(&User{
						Email:         fmt.Sprintf("%s@example.com", u),
						Username:      u,
						Password:      "",
						EmailVerifyAt: time.Now(),
					})
				}
				return nil
			},
		},
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
