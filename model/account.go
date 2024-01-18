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
	Profile   Profile `gorm:"foreignKey:AccountsId"`
}

type AccountResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}

type Profile struct {
	Introduction string `json:"introduction"`
	AccountsId   uint   `json:"AccountsId" gorm:"not null"`
}

type LoginInfo struct {
	gorm.Model
	Email    string `form:"email" binding:"required" gorm:"unique; not null"`
	Username string `form:"username" binding:"required" gorm:"not null"`
	Password string `form:"password" binding:"required" gorm:"not null"`
}

type SignupInfo struct {
	gorm.Model
	Firstname string `form:"firstname" binding:"required" gorm:"not null"`
	Lastname  string `form:"lastname" binding:"required" gorm:"not null"`
	Email     string `form:"email" binding:"required" gorm:"unique; not null"`

	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirmpassword" binding:"required"`
}
