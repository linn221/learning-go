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

func FindLoanByID(id uint) (*Loan, error) {
	// validate id
	err := (Loan{ID: id}).validateId()
	if err != nil {
		return &Loan{}, err
	}

	var result Loan
	err = DB.First(&result, id).Error
	return &result, err
}

func (input Loan) validateId() error {
	var count uint
	err := DB.Model(&Loan{}).Where("id = ?", input.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("loan not found")
	}
	return nil
}

func (input *Loan) Update() (*Loan, error) {
	// validation here
	err := input.validateId()
	if err != nil {
		return &Loan{}, err
	}

	err = DB.Model(&input).Updates(Loan{
		Name:   input.Name,
		Amount: input.Amount,
	}).Error
	return input, err
}

func (input *Loan) Delete() (*Loan, error) {
	err := input.validateId()
	if err != nil {
		return &Loan{}, err
	}

	err = DB.Delete(&input).Error
	return input, err
}
