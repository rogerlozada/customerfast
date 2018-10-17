package customerdomainmodel

import "gopkg.in/mgo.v2/bson"

//ICustomerUseCase customer interfaces
type ICustomerUseCase interface {
	GetByID(id string) (Customer, error)
	GetAll() ([]Customer, error)
	Add(customer Customer) (bson.ObjectId, error)
	Delete(id string) error
	Update(customer Customer) error
}
