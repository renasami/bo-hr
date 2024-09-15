package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/renasami/bo-hr/hr-service/domain/base"
	"github.com/renasami/bo-hr/hr-service/internal/db"
)

func CreateUserTitle(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close(context.Background())

	// リクエストボディの読み取り
	var userTitle UserTitle
	if err := json.NewDecoder(r.Body).Decode(&userTitle); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

}

func EditUserTitle() {}
