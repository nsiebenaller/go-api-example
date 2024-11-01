package models

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uuid.UUID `gorm:"primaryKey,default:uuid_generate_v4()" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updated_at"`
}

type MemberWithBooks struct {
	ID        uuid.UUID `gorm:"primaryKey,default:uuid_generate_v4()" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updated_at"`
	Books     []*Book   `gorm:"many2many:member_books;foreignKey:ID;joinForeignKey:MemberID" json:"books"`
}

type CreateMember struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type MemberBooks struct {
	MemberID string `json:"member_id"`
	BookID   string `json:"book_id"`
}
