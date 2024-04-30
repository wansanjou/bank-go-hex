package repository

import "gorm.io/gorm"

type accountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepositoryDB(db *gorm.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) CreateAccount(acc Account) (*Account , error) {
	err := r.db.Create(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (r accountRepositoryDB) GetAllAccount(customerID int) ([]Account , error) {
	accounts := []Account{}
	err := r.db.Where("customer_id = ?", customerID).Find(&accounts).Error
	if err != nil {
			return nil, err
	}
	return accounts, nil
}