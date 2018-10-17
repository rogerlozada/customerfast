package customerdomaininfraestructure

import (
	"errors"
	"os"

	"customerfast/api/src/domain.model/customer"
	"gopkg.in/mgo.v2/bson"
)

type (
	//CustomerRepository customer instances
	CustomerRepository struct{}
)

const tableName = "customer"

//NewCustomerRepository customer instances
func NewCustomerRepository() customerdomainmodel.ICustomerRepository {
	return &CustomerRepository{}
}

//GetByID return customer by ID
func (customerRepository CustomerRepository) GetByID(id string) (customerdomainmodel.Customer, error) {

	customer := customerdomainmodel.Customer{}
	if !bson.IsObjectIdHex(id) {
		return customer, errors.New("Customer not found")
	}

	idHex := bson.ObjectIdHex(id)

	mongoDBSession := getContext()
	defer mongoDBSession.Close()

	if err := mongoDBSession.DB(os.Getenv("MONGO_DATABASE")).C(tableName).FindId(idHex).One(&customer); err != nil {
		return customer, errors.New("Customer not found")
	}

	return customer, nil
}

//GetAll return all customers
func (customerRepository CustomerRepository) GetAll() ([]customerdomainmodel.Customer, error) {

	customer := []customerdomainmodel.Customer{}

	mongoDBSession := getContext()
	defer mongoDBSession.Close()

	if err := mongoDBSession.DB(os.Getenv("MONGO_DATABASE")).C(tableName).Find(nil).All(&customer); err != nil {
		return customer, errors.New("Customer(s) not found")
	}

	return customer, nil
}

//Add insert new customer in mongoDB
func (customerRepository CustomerRepository) Add(customer customerdomainmodel.Customer) (bson.ObjectId, error) {
	customer.ID = bson.NewObjectId()

	mongoDBSession := getContext()
	defer mongoDBSession.Close()

	err := mongoDBSession.DB(os.Getenv("MONGO_DATABASE")).C(tableName).Insert(customer)

	if err != nil {
		return customer.ID, err
	}

	return customer.ID, nil
}

//Delete delete a customer
func (customerRepository CustomerRepository) Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Customer not found")
	}

	idHex := bson.ObjectIdHex(id)

	mongoDBSession := getContext()
	defer mongoDBSession.Close()

	err := mongoDBSession.DB(os.Getenv("MONGO_DATABASE")).C(tableName).RemoveId(idHex)

	if err != nil {
		return err
	}

	return nil
}

//Update update a customer
func (customerRepository CustomerRepository) Update(customer customerdomainmodel.Customer) error {

	mongoDBSession := getContext()
	defer mongoDBSession.Close()

	err := mongoDBSession.DB(os.Getenv("MONGO_DATABASE")).C(tableName).UpdateId(customer.ID, customer)

	if err != nil {
		return errors.New("It's not possivel update customer")
	}

	return nil
}
