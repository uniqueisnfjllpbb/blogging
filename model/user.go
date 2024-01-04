package model

import (
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	//ID        uuid.UUID `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Firstname string  `json:"firstname" gorm:"not null"`
	Lastname  string  `json:"lastname" gorm:"not null"`
	Email     string  `json:"email" gorm:"type:varchar(120);unique_index; not null; unique"`
	Password  string  `json:"password" gorm:"not null"`
	Isactive  bool    `json:"isactive" gorm:"default:true"`
	Post      Post    `gorm:"foreignKey:AccountsId"`
	Profile   Profile `gorm:"foreignKey:AccountsId"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
type Profile struct {
	//Account      Accounts `json:"account" gorm:"foreignKey:ID; constraint:OnDelete:CASCADE"`
	Introduction string `json:"introduction"`
	AccountsId   uint   `json:"AccountsId" gorm:"not null"`
}
