package models

type Post struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"size:255; not null" json:"title"`
	Content    string `gorm:"text" json:"content"`
	CategoryID uint   `json:"category_id"`
	Category   *Category
}

func (input *Post) CreatePost() error {
	err := DB.Create(&input).Error
	return err
}

func (input *Post) UpdatePost() error {
	err := DB.Model(&input).Updates(Post{
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: input.CategoryID,
	}).Error
	return err
}

func (input *Post) DeletePost() error {
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
	err := DB.Preload("Category").First(&result, id).Error
	return result, err
}
