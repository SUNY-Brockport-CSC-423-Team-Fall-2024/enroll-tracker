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

type TeacherCreation struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Office      string `json:"office"`
}

type AdministratorCreation struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Office      string `json:"office"`
}

type StudentUpdate struct {
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	MajorID     *int       `json:"major_id,omitempty"`
}

type TeacherUpdate struct {
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Office      *string    `json:"office,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
}

type AdministratorUpdate struct {
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Office      *string    `json:"office,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
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

type Teacher struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	AuthID      int        `json:"auth_id"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	Office      string     `json:"office"`
	LastLogin   *time.Time `json:"last_login"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Administrator struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	AuthID      int        `json:"auth_id"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	Office      string     `json:"office"`
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
	Username       string    `json:"username"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshTokenID string    `json:"refresh_token_id"`
	IssuedAt       time.Time `json:"issued_at"`
	CreatedAt      time.Time `json:"created_at"`
	ExpiresAt      time.Time `json:"expires_at"`
	Revoked        bool      `json:"revoked"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

type TokenResponse struct {
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshTokenID string    `json:"refresh_token_id"`
	ExpiresIn      time.Time `json:"expires_in"`
}

type RouteAuth struct {
	MethodRoles []MethodRoles
}

type MethodRoles struct {
	Roles       []string `json:"role"`
	HTTPMethods []string `json:"http_methods"`
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
