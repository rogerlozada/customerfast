package customerapplication

import (
	"customerfast/api/src/customerContext/domain.model/customer"
	"gopkg.in/mgo.v2/bson"
)

type customerUseCase struct {
	customerRepository customerdomainmodel.ICustomerRepository
}

//NewCustomerUseCase customer use case instances
func NewCustomerUseCase(customerRepository customerdomainmodel.ICustomerRepository) customerdomainmodel.ICustomerUseCase {
	return &customerUseCase{customerRepository}
}

//GetByID return customer by ID
func (customerUseCase customerUseCase) GetByID(id string) (customerdomainmodel.Customer, error) {
	return customerUseCase.customerRepository.GetByID(id)
}

func (customerUseCase customerUseCase) GetAll() ([]customerdomainmodel.Customer, error) {
	return customerUseCase.customerRepository.GetAll()
}

//Add new customer
func (customerUseCase customerUseCase) Add(customer customerdomainmodel.Customer) (bson.ObjectId, error) {
	return customerUseCase.customerRepository.Add(customer)
}

//Delete a customer
func (customerUseCase customerUseCase) Delete(id string) error {
	return customerUseCase.customerRepository.Delete(id)
}

//Update a customer
func (customerUseCase customerUseCase) Update(customer customerdomainmodel.Customer) error {
	return customerUseCase.customerRepository.Update(customer)
}
