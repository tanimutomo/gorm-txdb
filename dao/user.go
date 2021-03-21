package dao

import (
	"github.com/tanimutomo/gorm-txdb/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{
		db: db,
	}
}

func (u User) Get(id int64) (*entity.User, error) {
	var e entity.User
	if err := u.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}
