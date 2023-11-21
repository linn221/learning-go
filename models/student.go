package models

type Student struct {
	ID   uint64 `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:256;not null;index" json:"name"`
	Age  uint   `gorm:"not null" json:"age"`
	Note string `gorm:"type:text" json:"note"`
}

// CREATE
func (input *Student) Store() error {
	// save db
	err := DB.Create(&input).Error
	if err != nil {
		return err
	}
	return nil
}

// READ
func IndexStudent() ([]Student, error) {
	var results []Student
	err := DB.Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

func GetStudentByID(id string) (*Student, error) {
	// assume input has id field present and valid
	var result Student
	err := DB.Find(&result, id).Error
	if err != nil {
		return &Student{}, err
	}
	return &result, nil
}

// UPDATE
func (input *Student) Update() error {
	// assume input has valid id
	// validate the other fields
	err := DB.Model(&input).Updates(Student{
		Name: input.Name,
		Age:  input.Age,
		Note: input.Note,
	}).Error
	return err
}

func (input *Student) Delete() error {
	err := DB.Delete(&input).Error
	return err
}
