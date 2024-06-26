package models

import(
	"gorm.io/gorm"
)

type Chat struct{
	gorm.Model
	Question string `gorm:"not null"`
	Answer string
	IsOpen bool `gorm:"default:false"`
	UserID uint 
	User User `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}