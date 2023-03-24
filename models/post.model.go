package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// this stuct will be transformed into the corresponding SQL table in the DB
type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"` // evoke the uuid_generate fn to gen a uuid for each record
	Title     string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Image     string    `gorm:"not null" json:"image,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePost struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
