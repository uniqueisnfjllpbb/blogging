package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	//ID        uuid.UUID `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Title string `json:"title" gorm:"not null"`
	Body  string `json:"post" gorm:"not null"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	//Account    Accounts `json:"account" gorm:"foreignKey:AccountsId; constraint:OnDelete:CASCADE"`
	AccountsId uint `json:"AccountsId" gorm:"not null"`
}

type Reply struct {
	ID        uuid.UUID `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Comment   string    `json:"comment" gorm:"not null"`
	Datetime  time.Time `json:"timestamp"`
	Account   Accounts  `json:"account" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
