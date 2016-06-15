package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

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

// BeforeCreate is an unexported type.
func (obj *Todo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
