package models

import "errors"

type Loan struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"size:256; index; not null"`
	Amount uint64 `gorm:"not null"`
}

func (input *Loan) Store() (*Loan, error) {
	// validation here
	err := DB.Create(&input).Error
	return input, err
}

func ListLoans() ([]Loan, error) {
	var results []Loan
	err := DB.Find(&results).Error
	return results, err
}

func FindLoanByID(id int) (*Loan, error) {
	// validate id
	var count uint8
	err := DB.Model(&Loan{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return &Loan{}, err
	}
	if count <= 0 {
		return &Loan{}, errors.New("Id does not exist")
	}

	var result Loan
	err = DB.First(&result, id).Error
	return &result, err
}

func (input *Loan) Update() (*Loan, error) {
	// validation here
	err := DB.Model(&input).Updates(Loan{
		Name:   input.Name,
		Amount: input.Amount,
	}).Error
	return input, err
}

func (input *Loan) Delete() (*Loan, error) {
	err := DB.Delete(&input).Error
	return input, err
}
