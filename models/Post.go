package models

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/linn221/go-blog/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Post struct {
	ID         uint      `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Title      string    `gorm:"size:255; not null" json:"title" validate:"required,min=3,max=255"`
	Slug       string    `gorm:"size:255; not null; unique" json:"slug" validate:"isdefault"`
	Content    string    `gorm:"text" json:"content" validate:"omitempty,min=5"`
	CategoryID uint      `json:"category_id" validate:"required,number,gte=0"`
	Category   *Category `json:"category,omitempty"`
	Tags       []Tag     `gorm:"many2many:post_tag;" json:"tags,omitempty"`
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

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	// create slug
	post.Slug = slug.Make(post.Title)
	return nil
}

func (input *Post) BeforeSave(tx *gorm.DB) error {
	input.Title = helpers.SanitizeStr(input.Title)
	input.Content = helpers.SanitizeStr(input.Content)
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

func GetPostBySlug(slug string) (Post, error) {
	var result Post
	// validate slug
	var count int64
	if err := DB.Model(&Post{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
		return result, err
	}
	if count <= 0 {
		return result, errors.New("slug does not exist")
	}
	err := DB.Where("slug = ?", slug).Preload(clause.Associations).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
