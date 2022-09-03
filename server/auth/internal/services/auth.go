package services

import "e-commerce/repositories"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repo *repositories.Repository) *Service {
		return &Service{}
}
