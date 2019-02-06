package router

import (
	"github.com/shyam-unnithan/go-restful/mapstore"

	"github.com/gorilla/mux"
)

//InitRoutes ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetCustomerRoutes(router, mapstore.NewMapStore())
	return router
}
