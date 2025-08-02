package service

import "user-auth-service/internal/repository"

type Authorization interface{}

type Users interface{}

type Admins interface{}

type StoreManager interface{}

type Service struct {
	Authorization
	Users
	Admins
	StoreManager
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
