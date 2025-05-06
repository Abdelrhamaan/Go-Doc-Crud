package models

import (
	"time"
)
type Document struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Title      string     `gorm:"type:varchar(100)" json:"title"`
	Author     string     `gorm:"type:varchar(100)" json:"author"`
	Content    string     `gorm:"type:text" json:"content"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
}