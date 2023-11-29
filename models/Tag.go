package models

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id" validate:"isdefault"`
	Name  string `gorm:"size=255; unique; not null" json:"name" validate:"required,min=3,max=255"`
	Posts []Post `gorm:"many2many:post_tag" validate:"isdefault"`
}
