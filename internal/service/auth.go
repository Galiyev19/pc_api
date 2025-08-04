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

func (s *Auth) Login(input models.InputRequest) (string, error) {
	// Implementation for login
	user,err := s.repo.Login(input.Email)
	if err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(input.Password)); err != nil {
		return 0,err
	
	}

	token, err := GenerateToken(input)
	if err != nil {
		return 0, err
	}

	return token,nil
}

func (s *Auth) SignUp(input models.SignUpRequest) (int, error) {
	// Implementation for sign up
	return 0, nil
}

func (s *Auth)GenerateToken(input models.SignUpRequest) (string, error) {
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


