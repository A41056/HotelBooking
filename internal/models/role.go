package models

import (
	"github.com/google/uuid"
)

// Role model
type Role struct {
	Base
	Name string `json:"name"`
}

// UserRole model to map users and roles
type UserRole struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}
