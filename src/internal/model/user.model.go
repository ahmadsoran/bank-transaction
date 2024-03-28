package model

import "gorm.io/gorm"

const StructUser = "User"

type User struct {
	Username string `json:"username" gorm:"unique" index:"idx_username"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model
}
