package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	fmt.Println(b.ID)
	return nil
}
