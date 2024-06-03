package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_const "main.go/internal/const"
	"main.go/internal/models"
	"main.go/internal/utils"
	"time"

	"gorm.io/gorm"
)

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

type userRepo struct {
	db     *gorm.DB
	hasher *utils.Hasher
}

func NewUserRepository(db *gorm.DB, hasher *utils.Hasher) UserRepository {
	return &userRepo{db, hasher}
}

func (ur *userRepo) Register(ctx context.Context, request *models.UserCreateRequest) error {
	salt, err := ur.hasher.RandomStr(16)
	if err != nil {
		return err
	}

	hashedPassword, err := ur.hasher.HashPassword(salt, request.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: hashedPassword,
		PasswordSalt: salt,
	}

	err = ur.db.Create(user).Error
	if err != nil {
		return err
	}

	defaultRole := &models.Role{Name: "customer"}
	err = ur.db.FirstOrCreate(defaultRole, defaultRole).Error
	if err != nil {
		return err
	}

	userRole := &models.UserRole{
		UserID: user.ID,
		RoleID: defaultRole.ID,
	}

	err = ur.db.Create(userRole).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) SeedData() error {
	// Check if the admin user already exists
	var adminUser models.User
	result := ur.db.Where("email = ?", "admin@example.com").First(&adminUser)
	if result.Error == nil {
		// Admin user already exists, nothing to do
		fmt.Println("Admin user already exists")
		return nil
	}

	// Create a new admin user
	salt, err := ur.hasher.RandomStr(16)
	if err != nil {
		return err
	}

	hashedPassword, err := ur.hasher.HashPassword(salt, "admin@123") // Set a default password for admin
	if err != nil {
		return err
	}

	adminUser = models.User{
		Username:     "admin",
		Email:        "admin@example.com",
		PasswordHash: hashedPassword,
		PasswordSalt: salt,
	}

	// Create the admin user record
	err = ur.db.Create(&adminUser).Error
	if err != nil {
		return err
	}

	// Fetch the admin role
	adminRole := &models.Role{Name: "admin"} // Assuming "admin" is the role name for admins
	err = ur.db.FirstOrCreate(adminRole, adminRole).Error
	if err != nil {
		return err
	}

	// Associate the admin role with the admin user
	userRole := &models.UserRole{
		UserID: adminUser.ID,
		RoleID: adminRole.ID,
	}
	err = ur.db.Create(userRole).Error
	if err != nil {
		return err
	}

	fmt.Println("Admin user created successfully")

	return nil
}

func (ur *userRepo) Login(ctx context.Context, username, password string) (*models.Token, error) {
	var user models.User
	err := ur.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	if !ur.hasher.CompareHashPassword(user.PasswordHash, user.PasswordSalt, password) {
		return nil, errors.New(_const.ErrIncorrectPassword)
	}

	tokenString, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	token := &models.Token{
		AccessToken: tokenString,
		ExpiresAt:   expiresAt,
	}

	return token, nil
}

// EditProfile updates the profile of a user in the database
func (ur *userRepo) EditProfile(ctx context.Context, userID uuid.UUID, user *models.User) error {
	var existingUser models.User
	if err := ur.db.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		return err
	}

	existingUser.Email = user.Email

	err := ur.db.Save(&existingUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	err := ur.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, userID uuid.UUID, user *models.User) error {
	var existingUser models.User
	if err := ur.db.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		return err
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email

	err := ur.db.Save(&existingUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	var user models.User
	err := ur.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
