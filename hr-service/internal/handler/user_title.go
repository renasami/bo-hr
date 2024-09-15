package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/renasami/bo-hr/hr-service/domain"
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
	var userTitle domain.UserTitle
	if err := json.NewDecoder(r.Body).Decode(&userTitle); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO user_title (tenant_name_id, title_id, title_name, deleted_at) VALUES ($1, $2, $3, null)"
	title_id, _ := uuid.NewV4()
	_, err = conn.Exec(context.Background(), query, userTitle.TenantNameID, title_id, userTitle.TitleName)
	if err != nil {
		http.Error(w, "Failed to update user title", http.StatusInternalServerError)
		return
	}

}

func EditUserTitle() {}
