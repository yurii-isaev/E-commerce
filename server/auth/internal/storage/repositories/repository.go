package repositories

import (
	"github.com/jmoiron/sqlx"

	"e-commerce/internal/domains"
)

type Authorization interface {
	CreateUser(user domains.User) (string, error)
	GetUser(username, password string) (domains.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
