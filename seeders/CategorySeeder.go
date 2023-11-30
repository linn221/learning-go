package seeders

import (
	"github.com/linn221/go-blog/models"
)

func SeedCategories() error {
	names := []string{
		"News",
		"Business",
		"Entertainment",
		"IT ",
		"Travel",
		"Funny",
	}
	var categories []models.Category
	for _, name := range names {
		categories = append(categories, models.Category{Name: name})
	}
	err := models.DB.Create(&categories).Error
	return err
}
