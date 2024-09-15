package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tenant struct {
	TenantNameID string    `json:"tenant_name_id" db:"tenant_name_id"`
	TenantUUID   uuid.UUID `json:"tenant_uuid" db:"tenant_uuid"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// TenantUser struct represents the tenant_user table
type TenantUser struct {
	TenantNameID string    `json:"tenant_name_id" db:"tenant_name_id"`
	TenantUserID uuid.UUID `json:"tenant_user_id" db:"tenant_user_id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
}

// Branches struct represents the branches table
type Branch struct {
	TenantNameID string     `json:"tenant_name_id" db:"tenant_name_id"`
	BranchID     uuid.UUID  `json:"branch_id" db:"branch_id"`
	BranchName   string     `json:"branch_name" db:"branch_name"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Nullable column
}

// UserTitle struct represents the user_title table
type UserTitle struct {
	TenantNameID string     `json:"tenant_name_id" db:"tenant_name_id"`
	TitleID      uuid.UUID  `json:"title_id" db:"title_id"`
	TitleName    string     `json:"title_name" db:"title_name"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Nullable column
}
