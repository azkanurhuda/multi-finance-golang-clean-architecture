package user

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u User) FindByToken(db *gorm.DB, token string) (*entity.User, error) {
	var user entity.User
	err := db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) FindByNIK(db *gorm.DB, nik string) (*entity.User, error) {
	var user entity.User
	err := db.Where("nik = ?", nik).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Create(db *gorm.DB, user *entity.User) error {
	return db.Create(user).Error
}

func (u User) UpdateTokenByNIK(db *gorm.DB, user *entity.User, nik string) error {
	return db.Model(user).Where("nik = ?", nik).Update("token", user.Token).Error
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}
