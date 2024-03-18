package Golang_Gorm

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserID      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}
