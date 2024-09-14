package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/renasami/bo-hr/hr-service/internal/db"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, name, position FROM employees")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Position)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
			http.Error(w, "Failed to parse results", http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
