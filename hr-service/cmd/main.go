package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renasami/bo-hr/hr-service/internal/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", handler.GetEmployees).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Hello</h1>")
	}).Methods("GET")

	log.Println("Starting HR service on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
