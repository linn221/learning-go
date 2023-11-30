package models

import (
	"errors"

	"github.com/linn221/go-blog/helpers"
	"gorm.io/gorm/clause"
)

type Post struct {
	ID         uint   `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Title      string `gorm:"size:255; not null" json:"title" validate:"required,min=3,max=255"`
	Content    string `gorm:"text" json:"content" validate:"omitempty,min=5"`
	CategoryID uint   `json:"category_id" validate:"required,number,gte=0"`
	Category   *Category
	Tags       []Tag `gorm:"many2many:post_tag"`
}

func (input Post) exists() error {
	var count int64
	if err := DB.Model(&Post{}).Where("id = ?", input.ID).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("Post does not exist")
	}
	return nil
}

func (input *Post) CreatePost() error {
	// checks if the category exists
	if err := (Category{ID: input.CategoryID}).exists(); err != nil {
		return err
	}

	err := DB.Create(&input).Error
	return err
}

func (input *Post) UpdatePost() error {
	if err := input.exists(); err != nil {
		return err
	}

	// checks if the category exists
	if err := (Category{ID: input.CategoryID}).exists(); err != nil {
		return err
	}

	err := DB.Model(&input).Updates(Post{
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: input.CategoryID,
	}).Error
	return err
}

func (input *Post) DeletePost() error {
	if err := input.exists(); err != nil {
		return err
	}
	err := DB.Delete(&input).Error
	return err
}

func GetAllPosts() ([]Post, error) {
	var results []Post
	err := DB.Find(&results).Error
	return results, err
}

func GetPostById(id string) (Post, error) {
	var result Post
	if err := (Post{ID: helpers.StrToUInt(id)}).exists(); err != nil {
		return result, err
	}
	err := DB.Preload(clause.Associations).First(&result, id).Error
	return result, err
}
