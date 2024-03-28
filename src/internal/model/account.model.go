package model

import (
	"gorm.io/gorm"
	"main.go/src/internal/constant/enum"
)

type Account struct {
	gorm.Model
	ID       uint          `json:"id"`
	Owner    string        `json:"owner" gorm:"not null, index:owner"`
	Currency enum.Currency `json:"currency" gorm:"not null"`
	Balance  int64         `json:"balance" gorm:"not null, default:0"`
	Freezed  bool          `json:"freezed" gorm:"default:false"`
}
