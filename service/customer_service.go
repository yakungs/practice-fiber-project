package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {

	return customerService{custRepo: custRepo}
}

func (c customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := c.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnExpectedError()
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (c customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := c.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnExpectedError()
	}
	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
