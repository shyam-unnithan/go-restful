package main

import (
	"log"
	"net/http"

	"github.com/shyam-unnithan/go-restful/router"
)

func main() {

	r := router.InitRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
