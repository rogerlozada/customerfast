package customerapplication

import (
	"encoding/json"
	"net/http"

	"customerfast/api/src/framework/common"

	"customerfast/api/src/customerContext/domain.model/customer"
)

//CustomerHandler struct of customer
type CustomerHandler struct {
	customerUseCase customerdomainmodel.ICustomerUseCase
}

//NewCustomerHandler create customer handler instances
func NewCustomerHandler(a customerdomainmodel.ICustomerUseCase) *CustomerHandler {
	return &CustomerHandler{a}
}

//GetByID return customer by ID
func (customerHandler CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	customer, err := customerHandler.customerUseCase.GetByID(id)
	if err != nil {
		common.ResponseError(err.Error(), common.HTTPInternalError, w)
		return
	}

	common.ResponseSuccess(customer, common.HTTPOk, w)
}

//GetAll return all customers
func (customerHandler CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	customer, err := customerHandler.customerUseCase.GetAll()
	if err != nil {
		common.ResponseError(err.Error(), common.HTTPInternalError, w)
		return
	}

	common.ResponseSuccess(customer, common.HTTPOk, w)
}

//Add a customer
func (customerHandler CustomerHandler) Add(w http.ResponseWriter, r *http.Request) {
	customer := customerdomainmodel.Customer{}

	json.NewDecoder(r.Body).Decode(&customer)

	_, err := customerHandler.customerUseCase.Add(customer)

	if err != nil {
		common.ResponseError(err.Error(), common.HTTPInternalError, w)
		return
	}

	common.ResponseSuccess("Customer add with success", common.HTTPOk, w)
}

//Delete a customer
func (customerHandler CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := customerHandler.customerUseCase.Delete(id)

	if err != nil {
		common.ResponseError(err.Error(), common.HTTPInternalError, w)
		return
	}

	common.ResponseSuccess("Customer delete with success!", common.HTTPOk, w)
}

//Update a customer
func (customerHandler CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	customer := customerdomainmodel.Customer{}

	json.NewDecoder(r.Body).Decode(&customer)

	err := customerHandler.customerUseCase.Update(customer)

	if err != nil {
		common.ResponseError(err.Error(), common.HTTPInternalError, w)
		return
	}

	common.ResponseSuccess("Customer updating with success", common.HTTPOk, w)
}
