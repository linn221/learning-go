package models

import (
	"errors"
	"time"

	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/utils/token"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string `gorm:"size:256; unique; not null" json:"name" validate:"required,min=4,max=256"`
	Password  string `gorm:"size:256; not null" json:"password" validate:"required,min=8,max=256"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (input *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := helpers.HashStr(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPassword)

	input.Name = helpers.SanitizeStr(input.Name)
	return nil
}

func (input *User) CreateUser() error {
	err := DB.Create(&input).Error
	return err
}

func (input *User) Login() (string, error) {
	var u User
	if err := DB.Model(&User{}).Where("name = ?", input.Name).Take(&u).Error; err != nil {
		return "", err
	}

	if err := helpers.VerifyPassword(input.Password, u.Password); err != nil {
		return "", err
	}

	tokenStr, err := token.GenerateAuthToken(u.ID, u.Name)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GetAllUsers() ([]User, error) {
	var result []User
	err := DB.Select("id", "name", "created_at", "updated_at").Find(&result).Error
	return result, err
}

func GetUserById(id uint) (User, error) {
	var count int64
	if err := DB.Model(&User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return User{}, err
	}
	if count <= 0 {
		return User{}, errors.New("User does not exist")
	}

	var user User
	err := DB.Omit("password").Find(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
