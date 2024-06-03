package domain

import "github.com/google/uuid"

type Role struct {
	Base
	Name string
}

type UserRole struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}
