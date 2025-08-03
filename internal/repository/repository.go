package repository

import "user-auth-service/internal/models"

type Authorization interface {
	Login(input models.InputRequest) (int, error)
	SignUp(input models.SignUpRequest) (int, error)
}

type Users interface {
	InsertUser(input *models.User) (int, error)
	UpdateUser(id int, input models.User) (int, error)
	DeleteUser(id int) error
	GetUserByID(id int) (*models.User, error)
	GetUserList(limit, offset int) (*[]models.User, error)
	RefreshPassword(id int, input models.User) (int, error)
}

type Repository struct {
	Authorization
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}
