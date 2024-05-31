package repositories

import (
	"context"
	"github.com/google/uuid"
	"main.go/internal/models"
)

// UserRepository defines methods for user-related operations
type UserRepository interface {
	SeedData() error
	Register(ctx context.Context, user *models.UserCreateRequest) error
	Login(ctx context.Context, username, password string) (*models.Token, error)
	EditProfile(ctx context.Context, userID uuid.UUID, user *models.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, userID uuid.UUID, user *models.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
