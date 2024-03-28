package model

import "gorm.io/gorm"

type Transaction struct {
	ID            uint    `json:"id"`
	FromAccountID uint    `json:"from_account_id" gorm:"index:from_account_id_idx"`
	ToAccountID   uint    `json:"to_account_id" gorm:"index:to_account_id_idx"`
	FromAccount   Account `json:"from_account" gorm:"not null;foreignKey:FromAccountID"`
	ToAccount     Account `json:"to_account" gorm:"not null;foreignKey:ToAccountID"`
	Amount        int64   `json:"amount" gorm:"not null;note:'must be positive'"`
	gorm.Model
}
