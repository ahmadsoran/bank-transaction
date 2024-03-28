package model

import "gorm.io/gorm"

type Entries struct {
	gorm.Model
	ID        uint    `json:"id"`
	AccountID uint    `json:"account_id" gorm:"index:account_id_idx"`
	Account   Account `json:"account" gorm:"not null;foreignKey:AccountID"`
	Amount    int     `json:"amount"`
}
