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
	var count int64
	if err := DB.Model(&Category{}).Where("name = ?", input.Name).Count(&count).Error; err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("category name already exists")
	}

	err := DB.Create(&input).Error

	return err
}

func (input *Category) UpdateCategory() error {
	// checks if category exists
	if err := input.exists(); err != nil {
		return err
	}

	// checks for duplicate
	var count int64
	if err := DB.Model(&Category{}).Not("id = ?", input.ID).Where("name = ?", input.Name).Count(&count).Error; err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("category name already exists")
	}

	err := DB.Model(&input).Updates(Category{
		Name: input.Name,
	}).Error
	return err
}

func (input *Category) DeleteCategory() error {
	// checks if category exists
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
	// checks if category exists
	if err := (Category{ID: helpers.StrToUInt(id)}).exists(); err != nil {
		return result, err
	}

	err := DB.Preload(clause.Associations).First(&result, id).Error
	return result, err
}
