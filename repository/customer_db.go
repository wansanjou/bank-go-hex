package repository

import (
	"gorm.io/gorm"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepositoryDB(db *gorm.DB) CustomerRepository {
	return customerRepositoryDB{db : db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error)  {
	customers := []Customer{}
	err := r.db.Find(&customers).Error
	if err != nil {
		return nil , err
	}

	return customers , nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error)	{
	customer := Customer{}
	result := r.db.First(&customer , id)
	if result.Error != nil {
		return nil , result.Error
	}

	return &customer , nil
}