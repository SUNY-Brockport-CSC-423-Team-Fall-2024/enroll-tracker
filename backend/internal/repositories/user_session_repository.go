package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"errors"
	"time"
)

type UserSessionRepository interface {
	CreateUserSession(userID int, username string, refreshToken string, refreshTokenID string, issuedAt time.Time, expiresAt time.Time) (models.UserSession, error)
	GetUserSession(refreshTokenID string) (models.UserSession, error)
	GetUserRole(username string) (string, error)
	RevokeUserSession(refreshTokenID string) error
}

type PostgresUserSessionRepository struct {
	db *sql.DB
}

func CreatePostgresUserSessionRepository(db *sql.DB) *PostgresUserSessionRepository {
	return &PostgresUserSessionRepository{db: db}
}

func (r *PostgresUserSessionRepository) CreateUserSession(userID int, username string, refreshToken string, refreshTokenID string, issuedAt time.Time, expiresAt time.Time) (models.UserSession, error) {
	var userSession models.UserSession

	query := `SELECT * FROM public.create_user_session($1,$2,$3,$4,$5,$6)`

	row := r.db.QueryRow(query, userID, username, refreshToken, refreshTokenID, issuedAt, expiresAt)

	err := row.Scan(&userSession.ID, &userSession.UserID, &userSession.Username, &userSession.RefreshToken, &userSession.RefreshTokenID, &userSession.IssuedAt, &userSession.ExpiresAt, &userSession.Revoked)

	return userSession, err
}

func (r *PostgresUserSessionRepository) GetUserSession(refreshTokenID string) (models.UserSession, error) {
	var userSession models.UserSession

	query := `SELECT * FROM public.get_user_session($1)`

	row := r.db.QueryRow(query, refreshTokenID)

	err := row.Scan(&userSession.ID, &userSession.UserID, &userSession.Username, &userSession.RefreshToken, &userSession.RefreshTokenID, &userSession.IssuedAt, &userSession.ExpiresAt, &userSession.Revoked)

	return userSession, err
}

func (r *PostgresUserSessionRepository) GetUserRole(username string) (string, error) {
	var role string

	query := `SELECT * FROM public.get_user_role($1)`

	row := r.db.QueryRow(query, username)

	err := row.Scan(&role)

	return role, err
}

func (r *PostgresUserSessionRepository) RevokeUserSession(refreshTokenID string) error {
	query := `SELECT * FROM public.revoke_user_session($1)`

	result, err := r.db.Exec(query, refreshTokenID)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numRows != 1 {
		return errors.New("Unable to revoke user session")
	}

	return nil
}
