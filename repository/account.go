package repository

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	CustomerID  uint  `db:"customer_id"`
	Customer    Customer
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string `db:"status"`
}

type AccountRepository interface {
	CreateAccount(Account) (*Account , error)
	GetAllAccount(int) ([]Account , error)
}