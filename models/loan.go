package models

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
	var result Loan
	err := DB.First(&result, id).Error
	return &result, err
}
