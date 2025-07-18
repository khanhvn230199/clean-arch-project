package service

import (
	"clean-arch-project/internal/domain/entity"
	"errors"
	"strings"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) ValidateUser(user *entity.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if !s.isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func (s *UserService) isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func (s *UserService) SanitizeName(name string) string {
	return strings.TrimSpace(name)
}
