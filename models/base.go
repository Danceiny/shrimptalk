package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (b Base) IsNil() bool {
	return uuid.Equal(b.ID, uuid.Nil)
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV1())
	return nil
}
