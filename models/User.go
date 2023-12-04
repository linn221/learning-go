package models

import (
	"time"

	"github.com/linn221/go-blog/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string `gorm:"size:256; unique; not null" json:"name" validate:"required,min=4,max=256"`
	Password  string `gorm:"size:256; not null" json:"password" validate:"required,min=8,max=256"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (input *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := hash(input.Password)
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

func (input *User) Login() error {

	return nil
}

func GetAllUsers() ([]User, error) {
	var result []User
	err := DB.Select("id", "name", "created_at", "updated_at").Find(&result).Error
	return result, err
}
