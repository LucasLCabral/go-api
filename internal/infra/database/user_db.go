package database

import (
	"errors"

	"github.com/LucasLCabral/go-api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	_, err := u.FindByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
        return u.DB.Create(user).Error
    }
	return err
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
