package models

import (
	"errors"

	"github.com/linn221/go-blog/helpers"
)

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Name  string `gorm:"size=255; unique; not null" json:"name" validate:"required,min=3,max=255"`
	Posts []Post `gorm:"many2many:post_tag; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"isdefault"`
}

func (input Tag) exists() error {
	var count int64
	err := DB.Model(&Tag{}).Where("id = ?", input.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("tag does not exist")
	}
	return nil
}

func (input *Tag) CreateTag() error {
	// check for uniqueness
	var count int64
	if err := DB.Model(&Tag{}).Where("name = ?", input.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Tag name already exists")
	}
	err := DB.Create(&input).Error
	return err
}

func (input *Tag) UpdateTag() error {

	// validate id
	if err := input.exists(); err != nil {
		return err
	}

	// check for uniqueness
	var count int64
	if err := DB.Model(&Tag{}).Not("id = ?", input.ID).Where("name = ?", input.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Tag name already exists")
	}
	err := DB.Model(&input).Updates(Tag{Name: input.Name}).Error
	return err
}

func (input *Tag) DeleteTag() error {
	// validate id
	if err := input.exists(); err != nil {
		return err
	}

	err := DB.Delete(&input).Error
	return err
}
func GetAllTags() ([]Tag, error) {
	var results []Tag
	err := DB.Find(&results).Error
	return results, err
}

func GetTagById(id string) (Tag, error) {
	var result Tag
	// validate id
	if err := (Tag{ID: helpers.StrToUInt(id)}).exists(); err != nil {
		return result, err
	}

	err := DB.Preload("Posts").Find(&result, id).Error
	return result, err
}
