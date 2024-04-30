package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{Name: "mock name 1", City: "mock city 1", ZipCode: "mock zipcode 1", Status: "1"},
		{Name: "mock name 2", City: "mock city 2", ZipCode: "mock zipcode 2", Status: "0"},
	}

	return customerRepositoryMock{customers: customers}
}

func (cm customerRepositoryMock) GetAll() ([]Customer, error) {
	return cm.customers, nil
}

func (cm customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range cm.customers {
		if int(customer.ID) == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}