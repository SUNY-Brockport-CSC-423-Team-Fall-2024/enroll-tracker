package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type StudentQueryParams struct {
	Username    *string   `json:"username"`
	FirstName   *string   `json:"first_name"`
	LastName    *string   `json:"last_name"`
	PhoneNumber *string   `json:"phone_number"`
	Email       *string   `json:"email"`
	Limit       *int      `json:"limit"`
	Offset      *int      `json:"offset"`
	Majors      *[]string `json:"majors"`
}

type TeacherQueryParams struct {
	Username    *string `json:"username"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Office      *string `json:"office"`
	Limit       *int    `json:"limit"`
	Offset      *int    `json:"offset"`
}

type AdministratorQueryParams struct {
	Username    *string `json:"username"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
	Office      *string `json:"office"`
	Limit       *int    `json:"limit"`
	Offset      *int    `json:"offset"`
}

type CourseQueryParams struct {
	Limit         *int    `json:"limit"`
	Offset        *int    `json:"offset"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	TeacherID     *int    `json:"teacher_id"`
	MaxEnrollment *int    `json:"max_enrollment"`
	MinEnrollment *int    `json:"min_enrollment"`
	MaxNumCredits *int    `json:"max_num_credits"`
	MinNumCredits *int    `json:"min_num_credits"`
	Status        *string `json:"status"`
}

type MajorQueryParams struct {
	Limit       *int    `json:"limit"`
	Offset      *int    `json:"offset"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

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
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Email       *string `json:"email,omitempty"`
}

type TeacherUpdate struct {
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Email       *string `json:"email,omitempty"`
	Office      *string `json:"office,omitempty"`
}

type AdministratorUpdate struct {
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Email       *string `json:"email,omitempty"`
	Office      *string `json:"office,omitempty"`
}

type Student struct {
	Username    string    `json:"username"`
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	AuthID      int       `json:"auth_id"`
	MajorID     *int      `json:"major_id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Teacher struct {
	Username    string    `json:"username"`
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	AuthID      int       `json:"auth_id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Office      string    `json:"office"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Administrator struct {
	Username    string    `json:"username"`
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	AuthID      int       `json:"auth_id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Office      string    `json:"office"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAuthentication struct {
	ID                int        `json:"id"`
	Username          string     `json:"username"`
	PasswordHash      string     `json:"password_hash"`
	LastLogin         *time.Time `json:"last_login"`
	LastPasswordReset *time.Time `json:"last_password_reset"`
	IsActive          bool       `json:"is_active"`
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

type CoursesStudent struct {
	StudentUsername string     `json:"student_username"`
	StudentID       int        `json:"student_id"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	AuthID          int        `json:"auth_id"`
	MajorID         *int       `json:"major_id"`
	PhoneNumber     string     `json:"phone_number"`
	Email           string     `json:"email"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	IsEnrolled      bool       `json:"is_enrolled"`
	EnrolledDate    time.Time  `json:"enrolled_date"`
	UnenrolledDate  *time.Time `json:"unenrolled_date"`
}

type StudentsCourse struct {
	CourseID          int        `json:"course_id"`
	CourseName        string     `json:"course_name"`
	CourseDescription string     `json:"course_description"`
	TeacherID         int        `json:"teacher_id"`
	MaxEnrollment     int        `json:"max_enrollment"`
	NumCredits        int        `json:"num_credits"`
	Status            string     `json:"status"`
	LastUpdated       time.Time  `json:"last_updated"`
	CreatedAt         time.Time  `json:"created_at"`
	IsEnrolled        bool       `json:"is_enrolled"`
	EnrolledDate      time.Time  `json:"enrolled_date"`
	UnenrolledDate    *time.Time `json:"unenrolled_date"`
}

type TeachersCourse struct {
	CourseID          int       `json:"course_id"`
	CourseName        string    `json:"course_name"`
	CourseDescription string    `json:"course_description"`
	TeacherID         int       `json:"teacher_id"`
	CurrentEnrollment int       `json:"current_enrollment"`
	MaxEnrollment     int       `json:"max_enrollment"`
	NumCredits        int       `json:"num_credits"`
	Status            string    `json:"status"`
	LastUpdated       time.Time `json:"last_updated"`
	CreatedAt         time.Time `json:"created_at"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Role   string `json:"role"`
	UserID int    `json:"user_id"`
	AuthID int    `json:"auth_id"`
}

type TokenResponse struct {
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshTokenID string    `json:"refresh_token_id"`
	ExpiresIn      time.Time `json:"expires_in"`
}

type Course struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	TeacherID     int       `json:"teacher_id"`
	MaxEnrollment int       `json:"max_enrollment"`
	NumCredits    int       `json:"num_credits"`
	Status        string    `json:"status"`
	LastUpdated   time.Time `json:"last_updated"`
	CreatedAt     time.Time `json:"created_at"`
}

type CourseCreation struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	TeacherID     int    `json:"teacher_id"`
	MaxEnrollment int    `json:"max_enrollment"`
	NumCredits    int    `json:"num_credits"`
}

type CourseUpdate struct {
	Description   *string `json:"description"`
	TeacherID     *int    `json:"teacher_id"`
	MaxEnrollment *int    `json:"max_enrollment"`
	NumCredits    *int    `json:"num_credits"`
	Status        *string `json:"status"`
}

type Major struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}

type MajorCreation struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MajorUpdate struct {
	Description *string `json:"description"`
	Status      *string `json:"status"`
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

var NoAffectedRows = errors.New("No affected rows")
