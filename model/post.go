package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title      string `json:"title" gorm:"not null"`
	Body       string `json:"post" gorm:"not null"`
	AccountsId uint   `json:"AccountsId" gorm:"not null"`
}

type Reply struct {
	gorm.Model
	Comment   string    `json:"comment" gorm:"not null"`
	Account   Accounts  `json:"account" gorm:"foreignKey:AccountsId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
