package main

import (
	"assignments/4/router"
	"log"
	"net/http"
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
