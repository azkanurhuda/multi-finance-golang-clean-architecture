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

func (u User) FindByID(db *gorm.DB, id string) (*entity.User, error) {
	var user entity.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Create(db *gorm.DB, user *entity.User) error {
	return db.Create(user).Error
}

func (u User) UpdateTokenByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Model(user).Where("email = ?", email).Updates(map[string]interface{}{
		"token":            user.Token,
		"token_expired_at": user.TokenExpiredAt,
	}).Error
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}
