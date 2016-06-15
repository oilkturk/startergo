package model

import "time"

// Todo is an unexported type.
type Todo struct {
	ID            string `gorm:"type:char(36);column:todo_id;primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	TodoTitle     string `gorm:"not null;type:varchar(255)"`
	CompletedDate time.Time
	Priority      int
	Color         string `gorm:"type:varchar(255)"`
}

// TableName is an unexported type.
func (Todo) TableName() string {
	return "starter_todo"
}
