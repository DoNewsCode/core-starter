package repositories

import (
	"github.com/nfangxu/core-skeleton/app/entities"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (r UserRepository) GetOne(id int) entities.User {
	var user entities.User
	r.db.First(&user, id)
	return user
}
