package main

import (
	"fmt"
	"go_curd/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// starts a new mux router
	r := mux.NewRouter()

	// route handlers, mapping our api routes with their handler functions
	// r.HandleFunc("/route", FunctionName).Methods("Method")

	// using handler.FunctionName because we are defineing out handler functions in go_curd/handlers

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server Running on http://localhost:8000")

	// starts the server on localhost port 8000 and 'r' refers to http.handler that routes the incoming requests to the server
	log.Fatal(http.ListenAndServe(":8000", r))

}
