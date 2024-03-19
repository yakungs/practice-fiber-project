package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{
			CustomerID:  1,
			Name:        "Kazuya",
			DateOfBirth: "1997-04-17",
			City:        "BKK",
			ZipCode:     "10520",
			Status:      1,
		},
	}
	return customerRepositoryMock{customers: customers}
}

func (m customerRepositoryMock) GetAll() ([]Customer, error) {

	return m.customers, nil
}

func (m customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range m.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("Customer not found")
}
