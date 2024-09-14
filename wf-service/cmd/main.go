package main

import (
	"github.com/gorilla/mux"
	"github.com/renasami/bo-hr/wf-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/workflows", handler.GetWorkflows).Methods("GET")

	log.Println("Starting WF service on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
