package model

import (
	"github.com/linn221/auth/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Username string `gorm:"size:255;unique" json:"username"`
	Password string `gorm:"size:255; not null" json:"password"`
}

func (input *User) CreateUser() error {
	err := DB.Create(&input).Error
	return err
}

func hash(plaintext string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)

}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (input *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := hash(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPassword)
	return nil
}

func GetAllUsers() ([]User, error) {
	var results []User
	err := DB.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (input User) Login() (string, error) {
	var u User
	if err := DB.Model(&User{}).Where("username = ?", input.Username).Take(&u).Error; err != nil {
		return "", err
	}
	if err := verifyPassword(input.Password, u.Password); err != nil {
		return "", err
	}
	tokenStr, err := utils.GenerateToken(u.ID, u.Username)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
