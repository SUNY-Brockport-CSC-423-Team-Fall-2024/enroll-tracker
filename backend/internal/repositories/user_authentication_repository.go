package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"fmt"
)

type UserAuthenticationRepository interface {
	CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error)
	GetUserAuthentication(username string) (models.UserAuthentication, error)
}

type PostgresUserAuthenticationRepository struct {
	db *sql.DB
}

func (r *PostgresUserAuthenticationRepository) CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := fmt.Sprintf(`SELECT username, password_hash FROM public.create_user_auth($1, $2)`)

	err := r.db.QueryRow(query, username, passwordHash).Scan(&userAuth.Username, &userAuth.PasswordHash)

	return userAuth, err
}

func (r *PostgresUserAuthenticationRepository) GetUserAuthentication(username string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := fmt.Sprintf(`SELECT username, password_hash FROM public.user_authentication WHERE username=$1`)

	err := r.db.QueryRow(query, username).Scan(&userAuth.Username, &userAuth.PasswordHash)

	return userAuth, err
}
