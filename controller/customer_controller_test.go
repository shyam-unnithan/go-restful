package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shyam-unnithan/go-restful/domain"
	"github.com/shyam-unnithan/go-restful/mocks"
)

var tcc TestCustomerController

type TestCustomerController struct {
	postPayload string
	customer    domain.Customer
}

func setup() {
	tcc = TestCustomerController{
		postPayload: fmt.Sprint(`{ "id": "1", "name": "Shyam Unnithan","email": "shyam.unnithan@anz.com"}`),
		customer: domain.Customer{
			ID:    "1",
			Name:  "Shyam Unnithan",
			Email: "shyam.unnithan@anz.com",
		},
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func compareCustomer(x, y interface{}) bool {
	if x == nil || y == nil {
		return false
	}
	customerX, ok := x.(domain.Customer)
	if !ok {
		return false
	}
	customerY, ok := y.(domain.Customer)
	if !ok {
		return false
	}
	if customerX.ID != customerY.ID ||
		customerX.Name != customerY.Name ||
		customerX.Email != customerY.Email {
		return false
	}
	return true
}

func TestPostCustomer(t *testing.T) {
	w := httptest.NewRecorder()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCustomerStore := mocks.NewMockCustomerStore(mockCtrl)
	mockCustomerStore.EXPECT().GetAll().Return(map[string]domain.Customer{}).Times(1)
	mockCustomerStore.EXPECT().Create(tcc.customer).Return(nil).Times(1)

	customerController := CustomerController{
		Store: mockCustomerStore,
	}

	request, _ := http.NewRequest("POST", "/customers", strings.NewReader(tcc.postPayload))
	request.Header.Set("Content-Type", "application/json")
	cust, status, _ := customerController.PostCustomer(w, request)

	if compareCustomer(tcc.customer, cust) {
		t.Errorf("Customer expected: %s, received %s", tcc.customer, cust)
	}
	if status != 201 {
		t.Errorf("HTTP status expected: 201, received: %s ", strconv.Itoa(status))
	}
}
