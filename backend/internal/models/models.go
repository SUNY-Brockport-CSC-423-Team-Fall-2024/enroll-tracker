package models

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type StudentCreation struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type StudentUpdate struct {
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	MajorID     *int       `json:"major_id,omitempty"`
}

type Student struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	AuthID      int        `json:"auth_id"`
	MajorID     *int       `json:"major_id"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	LastLogin   *time.Time `json:"last_login"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type UserAuthentication struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserSession struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	SessID         string    `json:"sess_id"`
	IssuedAt       time.Time `json:"issued_at"`
	ExpirationTime time.Time `json:"expiration_time"`
	NotBefore      time.Time `json:"not_before"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	SessID string `json:"sess_id,omitempty"`
}

// Minimal wrapper for http.ResponseWriter that allows status code to be captured for logging.
type ResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}

func (rw *ResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}
func WrapResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}
