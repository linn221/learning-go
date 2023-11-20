package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type AccountType string

const (
	Assets      AccountType = "Assets"
	Liabilities AccountType = "Liabilities"
	Income      AccountType = "Income"
	Expenses    AccountType = "Expenses"
	Equity      AccountType = "Equity"
)

type Account struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	Name      string      `gorm:"size:255;not null" json:"name" binding:"required"`
	Type      AccountType `gorm:"type:enum('Assets', 'Liabilities', 'Income', 'Expenses', 'Equity')" json:"type"`
	CreatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (input *Account) BeforeSave() error {

	//remove spaces
	input.Name = html.EscapeString(strings.TrimSpace(input.Name))

	return nil
}

func (input *Account) CreateAccount() (*Account, error) {

	var count int64

	// check for duplicate accounts
	err := DB.Model(&Account{}).Where("name = ?", input.Name).Count(&count).Error
	if err != nil {
		return &Account{}, err
	}
	if count > 0 {
		return &Account{}, errors.New("duplicate account")
	}

	// actual creation
	err = DB.Create(&input).Error
	if err != nil {
		return &Account{}, err
	}
	return input, nil
}

func (input *Account) UpdateAccount() (*Account, error) {

	var count int64

	// checks for existance
	err := DB.Model(&Account{}).Where("id = ?", input.ID).Count(&count).Error
	if err != nil {
		return &Account{}, err
	}
	if count <= 0 {
		return &Account{}, errors.New("account not found")
	}

	// checks for duplicate
	err = DB.Model(&Account{}).Not("id = ?", input.ID).Where("name = ?", input.Name).Count(&count).Error
	if err != nil {
		return &Account{}, err
	}
	if count > 0 {
		return &Account{}, errors.New("duplicate account")
	}

	// actual update code
	err = DB.Model(&input).Updates(Account{Name: input.Name, Type: input.Type}).Error
	if err != nil {
		return &Account{}, err
	}
	return input, nil
}

func (input *Account) DeleteAccount() (*Account, error) {

	var count int64

	// validation
	err := DB.Model(&Account{}).Where("id = ?", input.ID).Count(&count).Error
	if err != nil {
		return &Account{}, err
	}
	if count <= 0 {
		return &Account{}, errors.New("account not found")
	}

	// actual deletion
	err = DB.Delete(&input).Error
	if err != nil {
		return &Account{}, err
	}
	return input, nil
}

func GetAccount(id string) (Account, error) {

	var result Account
	if err := DB.First(&result, id).Error; err != nil {
		return result, errors.New("account not found")
	}
	return result, nil
}

func GetAllAccounts() ([]Account, error) {

	var results []Account
	if err := DB.Find(&results).Error; err != nil {
		return results, errors.New("no account")
	}
	return results, nil
}
