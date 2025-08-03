package service

import (
	"user-auth-service/internal/models"
	"user-auth-service/internal/repository"
)

type Auth struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) Login(input models.InputRequest) (int, error) {
	// Implementation for login
	return 0, nil
}

func (s *Auth) SignUp(input models.SignUpRequest) (int, error) {
	// Implementation for sign up
	return 0, nil
}

func (s *Auth) GenerateToken(userID int) (string, error) {
	// Implementation for generating token
	return "", nil
}

func (s *Auth) RefreshToken(token string) (string, error) {
	// Implementation for refreshing token
	return "", nil
}

func (s *Auth) ParseToken(token string) (int, error) {
	// Implementation for parsing token
	return 0, nil
}
