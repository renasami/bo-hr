package main

import (
	"github.com/gorilla/mux"
	"github.com/renasami/bo-hr/hr-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", handler.GetEmployees).Methods("GET")

	log.Println("Starting HR service on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
