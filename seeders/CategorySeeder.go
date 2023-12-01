package seeders

import (
	"github.com/linn221/go-blog/data"
	"github.com/linn221/go-blog/models"
)

func SeedCategories() error {
	var categories []models.Category
	for _, name := range data.CategoryNames {
		categories = append(categories, models.Category{Name: name})
	}
	err := models.DB.Create(&categories).Error
	return err
}
