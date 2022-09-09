package services

import (
	"e-commerce/internal/entities"
	"e-commerce/internal/storage/repositories"
)

type Authorization interface {
	Registration(user entities.User) (string, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

// Return common istance on service authorization inplementation.
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.AuthCrudRepository),
	}
}
