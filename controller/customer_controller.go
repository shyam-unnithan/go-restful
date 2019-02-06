package controller

import (
	"assignments/4/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

//CustomerController manage customer
type CustomerController struct {
	Store domain.CustomerStore
}

//PostCustomer ...
func (handler CustomerController) PostCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var cust cust
	err := json.NewDecoder(r.Body).Decode(&cust)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Unable to decode json request body")
	}
	id := strconv.Itoa(len(handler.Store.GetAll()) + 1)

	c := domain.Customer{
		ID:    id,
		Name:  cust.Name,
		Email: cust.Email,
	}
	err = handler.Store.Create(c)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed creating customer")
	}
	return cust, http.StatusCreated, nil
}

//GetCustomer ...
func (handler CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var custs []cust
	for _, v := range handler.Store.GetAll() {
		cust := cust{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
		custs = append(custs, cust)
	}
	return custs, http.StatusCreated, nil
}

//GetCustomerByID ...
func (handler CustomerController) GetCustomerByID(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	c, err := handler.Store.GetByID(id)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Customer not found")
	}
	cust := cust{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
	}

	return cust, http.StatusNoContent, nil
}

//PutCustomer ...
func (handler CustomerController) PutCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var cust cust
	err := json.NewDecoder(r.Body).Decode(&cust)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Unable to decode json request body")
	}
	vars := mux.Vars(r)
	k := vars["id"]

	c := domain.Customer{
		ID:    k,
		Name:  cust.Name,
		Email: cust.Email,
	}
	err = handler.Store.Update(c)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Failed updating customer")
	}
	return nil, http.StatusNoContent, nil
}

//DeleteCustomer ...
func (handler CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	k := vars["id"]
	c, err := handler.Store.GetByID(k)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Unable to find customer")
	}
	err = handler.Store.Delete(c)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Customer deletion failed")
	}
	return nil, http.StatusNoContent, nil
}
