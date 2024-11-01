package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `gorm:"primaryKey,default:uuid_generate_v4()" json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updated_at"`
}

type CreateBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
