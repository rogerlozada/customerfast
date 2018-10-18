package customerdomainmodel

import "gopkg.in/mgo.v2/bson"

//ICustomerRepository customer repository interfaces
type ICustomerRepository interface {
	GetByID(id string) (Customer, error)
	GetAll() ([]Customer, error)
	Add(customer Customer) (bson.ObjectId, error)
	Delete(id string) error
	Update(customer Customer) error
}
