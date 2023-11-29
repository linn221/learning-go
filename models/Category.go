package models

import (
	"errors"
	"time"

	"github.com/linn221/go-blog/helpers"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID        uint   `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Name      string `gorm:"size:256; unique; not null" json:"name" validate:"required,min=3"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (input Category) exists() error {
	var count int64
	err := DB.Model(&Category{}).Where("id = ?", input.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Category does not exist")
	}
	return nil
}

func (input *Category) CreateCategory() error {
	// checks for duplicate
	err := DB.Create(&input).Error

	return err
}

func (input *Category) UpdateCategory() error {
	if err := input.exists(); err != nil {
		return err
	}

	err := DB.Model(&input).Updates(Category{
		Name: input.Name,
	}).Error
	return err
}

func (input *Category) DeleteCategory() error {
	if err := input.exists(); err != nil {
		return err
	}

	err := DB.Delete(&input).Error
	return err
}

func GetAllCategories() ([]Category, error) {
	var results []Category
	err := DB.Find(&results).Error
	return results, err
}

func GetCategoryById(id string) (Category, error) {
	var result Category

	result.ID = helpers.StrToUInt(id)
	if err := result.exists(); err != nil {
		return result, err
	}

	err := DB.Preload(clause.Associations).First(&result, id).Error
	return result, err
}
