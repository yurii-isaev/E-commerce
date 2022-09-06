package repositories

import (
	"github.com/jmoiron/sqlx"

	"e-commerce/internal/domains"
)

type Authorization interface {
	CreateUser(user domains.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
