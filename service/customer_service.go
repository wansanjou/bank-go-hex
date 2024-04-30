package service

import (
	"wansanjou/errs"
	"wansanjou/logs"
	"wansanjou/repository"
)

type customerService struct {
	customer_Repo repository.CustomerRepository
}

func NewCustomerService(customer_Repo repository.CustomerRepository) customerService  {
	return customerService{customer_Repo : customer_Repo}
}

func (c customerService) GetAllCustomer() ([]CustomerResponse, error) {
	customers , err := c.customer_Repo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil , errs.NewUnexpectedError()
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			ID : customer.ID,
			Name : customer.Name,
			Status : customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses , nil
}

func (c customerService) GetCustomer(id int) (*CustomerResponse, error)  {
	customer, err := c.customer_Repo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Customer not found")
	}

	customerResponse := CustomerResponse{
		ID : customer.ID,
		Name : customer.Name,
		Status : customer.Status,
	} 

	return &customerResponse , nil
}