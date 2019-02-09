package domain

// To generate the mock implementation used
// mockgen -destination=mocks/mock_customer.go -package=mocks github.com/shyam-unnithan/go-restful/domain CustomerController

// Customer structure
type Customer struct {
	ID, Name, Email string
}

//CustomerStore to manage customer
type CustomerStore interface {

	// Create customer information
	Create(Customer) error

	//Update customer information
	Update(Customer) error

	// Delete customer information
	Delete(Customer) error

	// Retrieve customer information by ID
	GetByID(string) (Customer, error)

	// Retrieve all customers
	GetAll() map[string]Customer
}
