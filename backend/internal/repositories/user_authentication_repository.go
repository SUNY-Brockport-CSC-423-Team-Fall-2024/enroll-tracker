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
	UpdateLastLogin(userAuthID int) (bool, error)
	DeleteUserAuthentication(username string) (bool, error)
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

	err := r.db.QueryRow(query, username, passwordHash).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash, &userAuth.LastLogin, &userAuth.IsActive)

	return userAuth, err
}

func (r *PostgresUserAuthenticationRepository) GetUserAuthentication(username string) (*models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	query := `SELECT id, username, password_hash, last_login, is_active FROM public.userauthentication WHERE username=$1`

	err := r.db.QueryRow(query, username).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash, &userAuth.LastLogin, &userAuth.IsActive)

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

	err := r.db.QueryRow(query, userAuthID, newPasswordHash).Scan(&userAuth.ID, &userAuth.Username, &userAuth.PasswordHash, &userAuth.LastLogin, &userAuth.LastPasswordReset, &userAuth.IsActive)

	if err != nil {
		return nil, err
	}

	return &userAuth, nil
}

func (r *PostgresUserAuthenticationRepository) UpdateLastLogin(userAuthID int) (bool, error) {
	query := `SELECT * FROM public.update_last_login($1)`

	result, err := r.db.Exec(query, userAuthID)
	if err != nil {
		return false, err
	}
	if rowsAffected, err := result.RowsAffected(); err == nil {
		if rowsAffected == 0 {
			return false, errors.New("Unable to update last login for user")
		}
	}

	return true, nil
}

func (r *PostgresUserAuthenticationRepository) DeleteUserAuthentication(username string) (bool, error) {
	var success bool

	query := `SELECT * FROM public.delete_user_auth($1)`

	err := r.db.QueryRow(query, username).Scan(&success)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}
