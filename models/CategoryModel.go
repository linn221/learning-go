package models

import (
	"time"

	"gorm.io/gorm/clause"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:256; unique; not null" json:"name"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (input *Category) CreateCategory() error {
	err := DB.Create(&input).Error

	return err
}

func (input *Category) UpdateCategory() error {
	err := DB.Model(&input).Updates(Category{
		Name: input.Name,
	}).Error
	return err
}

func (input *Category) DeleteCategory() error {
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
	err := DB.Preload(clause.Associations).First(&result, id).Error
	return result, err
}
