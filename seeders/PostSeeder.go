package seeders

import (
	"github.com/linn221/go-blog/factory"
	"github.com/linn221/go-blog/models"
)

func SeedPosts(count int) error {
	posts := factory.MakePosts(count)
	err := models.DB.Create(&posts).Error
	return err
}
