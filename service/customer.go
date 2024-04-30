package service

type CustomerResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type CustomerService interface {
	GetAllCustomer() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}