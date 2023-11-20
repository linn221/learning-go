package models

type Note struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	Title     string      `gorm:`
	Name      string      `gorm:"size:255;not null" json:"name" binding:"required"`
	Type      AccountType `gorm:"type:enum('Assets', 'Liabilities', 'Income', 'Expenses', 'Equity')" json:"type"`
	CreatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
