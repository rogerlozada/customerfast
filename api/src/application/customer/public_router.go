package customerapplication

import (
	"customerfast/api/src/domain.model/infraestructure"
	"github.com/gorilla/mux"
)

//GetCustomerRouter return customers router
func GetCustomerRouter(routerPrefix *mux.Router) *mux.Router {

	customerRepository := customerdomaininfraestructure.NewCustomerRepository()
	customerUseCase := NewCustomerUseCase(customerRepository)
	customerHandler := NewCustomerHandler(customerUseCase)

	publicRouter := mux.NewRouter()

	publicRouter.HandleFunc("/api/customer", customerHandler.GetByID).Methods("GET")
	publicRouter.HandleFunc("/api/allcustomer", customerHandler.GetAll).Methods("GET")
	publicRouter.HandleFunc("/api/customer", customerHandler.Add).Methods("POST")
	publicRouter.HandleFunc("/api/customer", customerHandler.Update).Methods("PUT")
	publicRouter.HandleFunc("/api/customer", customerHandler.Delete).Methods("DELETE")

	return publicRouter
}
