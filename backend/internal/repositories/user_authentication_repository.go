package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"errors"
)

type UserAuthenticationRepository interface {
	CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error)
	GetUserAuthentication(username string) (*models.UserAuthentication, error)
	ChangePassword(userAuthID int, newPasswordHash string) (*models.UserAuthentication, error)
}

type PostgresUserAuthenticationRepository struct {
	db *sql.DB
}

func CreatePostgresUserAuthenticationRepository(db *sql.DB) *PostgresUserAuthenticationRepository {
	return &PostgresUserAuthenticationRepository{db: db}
}

func (r *PostgresUserAuthenticationRepository) CreateUserAuthentication(username string, passwordHash string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := `SELECT * FROM public.create_user_auth($1, $2)`

	err := r.db.QueryRow(query, username, passwordHash).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash)

	return userAuth, err
}

func (r *PostgresUserAuthenticationRepository) GetUserAuthentication(username string) (*models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := `SELECT id, username, password_hash FROM public.userauthentication WHERE username=$1`

	err := r.db.QueryRow(query, username).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &userAuth, nil
}

func (r *PostgresUserAuthenticationRepository) ChangePassword(userAuthID int, newPasswordHash string) (*models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := `SELECT * FROM public.change_user_password($1,$2)`

	err := r.db.QueryRow(query, userAuthID, newPasswordHash).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash, &userAuth.LastPasswordReset)

	if err != nil {
		return nil, err
	}

	return &userAuth, nil
}
