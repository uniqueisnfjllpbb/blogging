package model

import (
	"github.com/oklog/ulid"
	"time"
)

type User struct {
	ID        ulid.ULID `gorm:"primaryKey;size:255;"`
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
