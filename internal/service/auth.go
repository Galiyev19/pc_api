package service

import (
	"time"
	"user-auth-service/internal/models"
	"user-auth-service/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) Login(input models.InputRequest) (models.TokenClaims, error) {
	// Implementation for login
	user, err := s.repo.Login(input.Email)
	if err != nil {
		return models.TokenClaims{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(input.Password)); err != nil {
		return models.TokenClaims{}, err

	}

	token, err := s.GenerateToken(input)
	if err != nil {
		return models.TokenClaims{}, err
	}

	return token, nil
}

func (s *Auth) SignUp(input models.InputRequest) (int, error) {
	var user models.User

	user.Email = input.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.HashPassword = string(hashedPassword)
	user.Role = "user" // Default role
	user.IsActive = true
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = user.CreatedAt

	id, err := s.repo.SignUp(user)
	if err != nil {
		return 0, err
	}
	// Implementation for sign up
	return id, nil
}

func (s *Auth) GenerateToken(input models.InputRequest) (models.TokenClaims, error) {
	claims := models.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// Implementation for generating token
	return claims, nil
}

func (s *Auth) RefreshToken(token string) (string, error) {
	// Implementation for refreshing token
	return "", nil
}

func (s *Auth) ParseToken(token string) (int, error) {
	// Implementation for parsing token
	return 0, nil
}
