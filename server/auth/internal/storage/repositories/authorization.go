package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"e-commerce/internal/domains"
	. "e-commerce/internal/storage/configs"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domains.User) (string, error) {
	var id string

	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3)`, UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return "error", err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (domains.User, error) {
	var user domains.User

	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, UsersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
