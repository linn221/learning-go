package models

import "time"

type Note struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (input *Note) store() (*Note, error) {
	// validation

	err := DB.Create(&input).Error
	if err != nil {
		return &Note{}, err
	}
	return input, nil
}

func (input *Note) update() (*Note, error) {
	var note Note
	DB.Find(&note, input.ID)
	DB.Model(&note).Updates(Note{
		Title: input.Title,
		Body:  input.Body,
	})
	return &note, nil
}

func GetNote(id string) (Note, error) {
	var note Note
	DB.First(&note, id)
	return note, nil
}

func ListNotes() ([]Note, error) {
	var notes []Note
	DB.Find(&notes)

}
