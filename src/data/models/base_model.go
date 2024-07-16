package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id         uint      `gorm:"primary_key" `
	CreatedAt  time.Time `gorm:"type:timestamp with time zone;not null" `
	ModifiedAt time.Time `gorm:"type:timestamp with time zone;null" `
	DeletedAt  time.Time `gorm:"type:timestamp with time zone;null" `

	CreatedBy  int `gorm:"not null" `
	ModifiedBy int `gorm:"null" `
	DeletedBy  int `gorm:"null" `
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.ModifiedAt = time.Now().UTC()
	m.ModifiedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if value != nil {
		userId = int(value.(float64))
	}
	m.DeletedAt = time.Now().UTC()
	m.DeletedBy = userId
	return
}
