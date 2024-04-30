package service

import (
	"net/http"
	"strings"
	"time"
	"wansanjou/errs"
	"wansanjou/logs"
	"wansanjou/repository"

	"github.com/gofiber/fiber/v2/log"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService  {
	return accountService{accRepo: accRepo}
}

func (as accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error)  {
	//Validate input 
	if request.Amount < 5000 {
		return nil , errs.AppError{
			Code: http.StatusBadRequest,
			Message: "Amount at least 5,000",
		}
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil , errs.AppError{
			Code: http.StatusBadRequest,
			Message: "Account type should be saving or checking",
		}
	}

	account := repository.Account{
		CustomerID: uint(customerID),
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount: request.Amount,
		Status: "1",
	} 

	newAcc , err := as.accRepo.CreateAccount(account)
	if err != nil {
		logs.Error(err)
		return nil , errs.NewUnexpectedError()

	}

	response := AccountResponse{
		OpeningDate:  newAcc.OpeningDate,
		AccountType:  newAcc.AccountType,
		Amount: 			newAcc.Amount,
		Status: 			newAcc.Status,
	}

	return &response , nil

}

func (as accountService) GetAccounts(customerID int) ([]AccountResponse, error)  {
	accounts , err := as.accRepo.GetAllAccount(customerID)
	if err != nil  {
		log.Error(err)
		return nil , errs.NewUnexpectedError()
	}

	response := []AccountResponse{}
	for _ , accounts := range accounts {
		response = append(response, AccountResponse{
			OpeningDate:  accounts.OpeningDate,
			AccountType: accounts.AccountType,
			Amount: accounts.Amount,
			Status: accounts.Status,
		})
	}
	return response , nil
}
