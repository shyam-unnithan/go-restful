package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shyam-unnithan/go-restful/domain"
	"github.com/shyam-unnithan/go-restful/mocks"
)

func TestPostCustomer(t *testing.T) {
	w := httptest.NewRecorder()
	mockCtrl := gomock.NewController(t)
	cjson := fmt.Sprint(`{ "id": "1", "name": "Shyam Unnithan","email": "shyam.unnithan@anz.com"}`)
	c := domain.Customer{
		ID:    "1",
		Name:  "Shyam Unnithan",
		Email: "shyam.unnithan@anz.com",
	}
	defer mockCtrl.Finish()
	mockCustomerStore := mocks.NewMockCustomerStore(mockCtrl)
	mockCustomerStore.EXPECT().GetAll().Return(map[string]domain.Customer{}).Times(1)
	mockCustomerStore.EXPECT().Create(c).Return(nil).Times(1)

	customerController := CustomerController{
		Store: mockCustomerStore,
	}

	request, _ := http.NewRequest("POST", "/customers", strings.NewReader(cjson))
	request.Header.Set("Content-Type", "application/json")
	cust, status, _ := customerController.PostCustomer(w, request)

	if reflect.DeepEqual(c, cust) {
		t.Errorf("Customer expected: %s, received %s", c, cust)
	}
	if status != 201 {
		t.Errorf("HTTP status expected: 201, received: %s ", strconv.Itoa(status))
	}
}
