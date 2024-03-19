package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(r repository.AccountRepository) AccountService {

	return accountService{accRepo: r}
}

func (s accountService) NewAccount(id int, request NewAccountRequest) (*AccountResponse, error) {

	//validate

	if request.Amount < 500 {
		return nil, errs.NewValidationError("Amount at least 500")
	}

	if strings.ToLower(request.AccountType) != "saving" || strings.ToLower(request.AccountType) != "cheking" {
		return nil, errs.NewValidationError("Account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  id,
		OpenDate:    time.Now().Format("2006-01-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.CreateAccount(account)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnExpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpenDate:    newAcc.OpenDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (s accountService) GetAccount(id int) ([]AccountResponse, error) {

	accounts, err := s.accRepo.GetAll(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnExpectedError()
	}

	response := []AccountResponse{}
	for _, account := range accounts {
		response = append(response, AccountResponse{
			AccountID:   account.AccountID,
			OpenDate:    account.OpenDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return response, nil
}
