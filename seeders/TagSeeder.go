package seeders

import (
	"github.com/linn221/go-blog/data"
	"github.com/linn221/go-blog/models"
)

func SeedTags() error {
	var tags []models.Tag
	for _, tagName := range data.TagNames {
		tags = append(tags, models.Tag{Name: tagName})
	}

	err := models.DB.Create(&tags).Error
	return err
}
