package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"main.go/internal/const"
	"main.go/internal/domain"
	"main.go/internal/repositories"
)

type UserService interface {
	Register(ctx context.Context, user *domain.UserCreateRequest) error
	Login(ctx context.Context, username, password string) (*domain.Token, error)
	EditProfile(ctx context.Context, userID uuid.UUID, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, userID uuid.UUID, user *domain.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, user *domain.UserCreateRequest) error {
	if err := validateUserCreateRequest(user); err != nil {
		return err
	}
	return s.userRepo.Register(ctx, user)
}

func (s *userService) Login(ctx context.Context, username, password string) (*domain.Token, error) {
	if len(password) < 6 {
		return nil, errors.New(_const.ErrPasswordTooShort)
	}
	return s.userRepo.Login(ctx, username, password)
}

func (s *userService) EditProfile(ctx context.Context, userID uuid.UUID, user *domain.User) error {
	if userID == uuid.Nil {
		return errors.New(_const.ErrInvalidUserID)
	}
	//if err := validateUser(user); err != nil {
	//	return err
	//}
	return s.userRepo.EditProfile(ctx, userID, user)
}

func (s *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New(_const.ErrInvalidUserID)
	}
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	if err := validateUser(user); err != nil {
		return err
	}
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, userID uuid.UUID, user *domain.User) error {
	if userID == uuid.Nil {
		return errors.New(_const.ErrInvalidUserID)
	}
	if err := validateUser(user); err != nil {
		return err
	}
	return s.userRepo.UpdateUser(ctx, userID, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if userID == uuid.Nil {
		return errors.New(_const.ErrInvalidUserID)
	}
	return s.userRepo.DeleteUser(ctx, userID)
}
