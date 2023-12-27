package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
type Profile struct {
	User         User   `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Introduction string `json:"instroduction"`
}
