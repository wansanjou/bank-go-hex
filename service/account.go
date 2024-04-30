package service

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	CustomerID  uint    `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(int) ([]AccountResponse, error)
}