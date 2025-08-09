package service

import (
	"user-auth-service/internal/models"
	"user-auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) Login(input models.InputRequest) (string, error) {
	// Implementation for login
	user, err := s.repo.Login(input.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(input.Password)); err != nil {
		return "", err

	}

	token, err := s.GenerateToken(input)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Auth) SignUp(input models.InputRequest) (int, error) {
	// Implementation for sign up
	return 0, nil
}

func (s *Auth) GenerateToken(input models.InputRequest) (string, error) {
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
