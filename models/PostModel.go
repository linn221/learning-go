package models

type Post struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"size:255; not null" json:"title"`
	Content    string `gorm:"text" json:"content"`
	CategoryID uint
	Category   *Category
}
