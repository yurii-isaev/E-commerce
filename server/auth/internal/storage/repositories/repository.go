package repositories

import (
	"github.com/jmoiron/sqlx"

	"e-commerce/internal/entities"
)

type AuthCrudRepository interface {
	CreateUser(user entities.User) (int, error)
	ReadUser(username, password string) (entities.User, error)
}

type Repository struct {
	AuthCrudRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthCrudRepository: NewAuthPostgres(db),
	}
}
