package services

import (
	"e-commerce/internal/entities"
	"e-commerce/internal/storage/repositories"
)

type Authentication interface {
	Registration(user entities.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authentication
}

// Return common istance on service authorization inplementation.
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repo.AuthCrudRepository),
	}
}
