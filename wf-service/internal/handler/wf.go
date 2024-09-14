package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/renasami/bo-hr/wf-service/internal/db"
)

type Workflow struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetWorkflows(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, title, status FROM workflows")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var workflows []Workflow
	for rows.Next() {
		var wf Workflow
		err = rows.Scan(&wf.ID, &wf.Title, &wf.Status)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
			http.Error(w, "Failed to parse results", http.StatusInternalServerError)
			return
		}
		workflows = append(workflows, wf)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflows)
}
