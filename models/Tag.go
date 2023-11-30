package models

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Name  string `gorm:"size=255; unique; not null" json:"name" validate:"required,min=3,max=255"`
	Posts []Post `gorm:"many2many:post_tag" validate:"isdefault"`
}

func (input *Tag) CreateTag() error {
	err := DB.Create(&input).Error
	return err
}

func (input *Tag) UpdateTag() error {
	err := DB.Model(&input).Updates(Tag{Name: input.Name}).Error
	return err
}

func (input *Tag) DeleteTag() error {
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
	err := DB.Preload("Posts").Find(&result, id).Error
	return result, err
}
