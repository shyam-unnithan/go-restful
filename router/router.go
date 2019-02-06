package router

import (
	"assignments/4/mapstore"

	"github.com/gorilla/mux"
)

//InitRoutes ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetCustomerRoutes(router, mapstore.NewMapStore())
	return router
}
