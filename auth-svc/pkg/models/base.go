package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model

	Id string `gorm:"type:uuid;primary_key;"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	base.Id = uuid.New().String()
	return
}
