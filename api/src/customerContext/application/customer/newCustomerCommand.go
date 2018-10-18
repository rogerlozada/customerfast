package customerapplication

import "customerfast/api/src/customerContext/domain.model/customer"

type (
	//NewCustomerCommand instances
	NewCustomerCommand struct {
		Name string `json:"name" bson:"name"`
	}
)

//ToCustomer create customers intance
func (command NewCustomerCommand) ToCustomer() *customerdomainmodel.Customer {
	customer := new(customerdomainmodel.Customer)
	customer.Name = command.Name
	return customer
}
