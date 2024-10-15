package models

import (
	"net/http"
	"time"
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
	Username    string     `json:"username,omitempty"`
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
