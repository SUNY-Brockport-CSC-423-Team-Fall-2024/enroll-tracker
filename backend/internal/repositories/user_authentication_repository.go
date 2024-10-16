package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"errors"
	"fmt"
)

type UserAuthenticationRepository interface {
	CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error)
	GetUserAuthentication(username string) (models.UserAuthentication, error)
}

type PostgresUserAuthenticationRepository struct {
	db *sql.DB
}

func CreatePostgresUserAuthenticationRepository(db *sql.DB) *PostgresUserAuthenticationRepository {
	return &PostgresUserAuthenticationRepository{db: db}
}

func (r *PostgresUserAuthenticationRepository) CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := fmt.Sprintf(`SELECT * FROM public.create_user_auth($1, $2)`)

	err := r.db.QueryRow(query, username, passwordHash).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash)

	return userAuth, err
}

func (r *PostgresUserAuthenticationRepository) GetUserAuthentication(username string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := fmt.Sprintf(`SELECT id, username, password_hash FROM public.userauthentication WHERE username=$1`)

	err := r.db.QueryRow(query, username).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userAuth, nil
		}
		return userAuth, err
	}

	return userAuth, nil
}
