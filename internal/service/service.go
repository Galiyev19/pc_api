package service

import (
	"user-auth-service/internal/models"
	"user-auth-service/internal/repository"
)

type Authorization interface {
	Login(input models.InputRequest) (string, error)
	SignUp(input models.InputRequest) (int, error)
	GenerateToken(input models.InputRequest) (string, error)
	RefreshToken(token string) (string, error)
	ParseToken(token string) (int, error)
}

type Users interface {
	CreateUser(input *models.User) (int, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, input models.User) (int, error)
	DeleteUser(id int) error
	RefreshPassword(id int, input models.InputRequest) (int, error)
	GetUserList(limit, offset int) (*[]models.User, error)
}

type StoreManager interface{}

type Service struct {
	Authorization
	Users
	StoreManager
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
