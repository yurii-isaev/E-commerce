package services

import (
	"e-commerce/internal/domains"
	"e-commerce/internal/storage/repositories"
)

type Authorization interface {
	CreateUser(user domains.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
