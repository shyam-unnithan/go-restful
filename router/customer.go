package router

import (
	"assignments/4/controller"
	"assignments/4/domain"

	"github.com/gorilla/mux"
)

//SetCustomerRoutes ....
func SetCustomerRoutes(router *mux.Router, store domain.CustomerStore) *mux.Router {
	customerStore := store
	customerController := controller.CustomerController{Store: customerStore}
	router.Handle("/api/customers", controller.ResponseHandler(customerController.PostCustomer)).Methods("POST")
	router.Handle("/api/customers", controller.ResponseHandler(customerController.GetCustomer)).Methods("GET")
	router.Handle("/api/customers/{id}", controller.ResponseHandler(customerController.GetCustomerByID)).Methods("GET")
	router.Handle("/api/customers/{id}", controller.ResponseHandler(customerController.PutCustomer)).Methods("PUT")
	router.Handle("/api/customers/{id}", controller.ResponseHandler(customerController.DeleteCustomer)).Methods("DELETE")
	return router
}
