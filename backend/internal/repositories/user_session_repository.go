package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"fmt"
	"time"
)

type UserSessionRepository interface {
	CreateUserSession(user_id int, sess_id string, issued_at time.Time, expiration_time time.Time, not_before time.Time) (models.UserSession, error)
	GetUserSession(sess_id string) (models.UserSession, error)
}

type PostgresUserSessionRepository struct {
	db *sql.DB
}

func CreatePostgresUserSessionRepository(db *sql.DB) *PostgresUserSessionRepository {
	return &PostgresUserSessionRepository{db: db}
}

func (r *PostgresUserSessionRepository) CreateUserSession(user_id int, sess_id string, issued_at time.Time, expiration_time time.Time, not_before time.Time) (models.UserSession, error) {
	var userSession models.UserSession

	query := fmt.Sprintf(`SELECT * FROM public.create_user_session($1,$2,$3,$4,$5)`)

	row := r.db.QueryRow(query, user_id, sess_id, issued_at, expiration_time, not_before)

	err := row.Scan(&userSession.ID, &userSession.UserID, &userSession.SessID, &userSession.IssuedAt, &userSession.ExpirationTime, &userSession.NotBefore)

	return userSession, err
}

func (r *PostgresUserSessionRepository) GetUserSession(sess_id string) (models.UserSession, error) {
	var userSession models.UserSession

	query := `SELECT * FROM public.get_user_session($1)`

	row := r.db.QueryRow(query, sess_id)

	err := row.Scan(&userSession.ID, &userSession.UserID, &userSession.SessID, &userSession.IssuedAt, &userSession.ExpirationTime, &userSession.NotBefore)

	return userSession, err
}
