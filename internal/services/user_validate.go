package services

import (
	"errors"
	"main.go/internal/const"
	"main.go/internal/models"
	"regexp"
	"strings"
)

func validateUserCreateRequest(user *models.UserCreateRequest) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New(_const.ErrUserNameRequire)
	}
	if !isValidEmail(user.Email) {
		return errors.New(_const.ErrInvalidEmailFormat)
	}
	if len(user.Password) < 6 {
		return errors.New(_const.ErrPasswordTooShort)
	}
	return nil
}

func validateUser(user *models.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New(_const.ErrUserNameRequire)
	}
	if !isValidEmail(user.Email) {
		return errors.New(_const.ErrPasswordTooShort)
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
