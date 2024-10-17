package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"fmt"

	_ "golang.org/x/tools/go/analysis/passes/ifaceassert"
)

type StudentRepository interface {
	CreateStudent(firstName string, lastName string, authId int, phoneNumber string, email string) (models.Student, error)
	GetStudent(username string) (models.Student, error)
	UpdateStudent(username string, studentUpdates models.StudentUpdate) (models.Student, error)
	//DeleteStudent(username string) (bool, error)
}

type PostgresStudentRepository struct {
	db *sql.DB
}

func CreatePostgresStudentRepository(db *sql.DB) *PostgresStudentRepository {
	return &PostgresStudentRepository{db: db}
}

func (r *PostgresStudentRepository) CreateStudent(firstName string, lastName string, authId int, phoneNumber string, email string) (models.Student, error) {
	//create student
	var student models.Student

	//create query
	query := fmt.Sprintf(`SELECT * FROM public.create_student($1, $2, $3, $4, $5)`)

	//execute function, instantiate student, close connection
	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email).Scan(&student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}

func (r *PostgresStudentRepository) GetStudent(username string) (models.Student, error) {
	//create student
	var student models.Student

	//create query
	query := fmt.Sprintf(`SELECT * FROM public.get_student($1)`)

	//execute function
	row := r.db.QueryRow(query, username)

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.LastLogin, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}

func (r *PostgresStudentRepository) UpdateStudent(username string, studentUpdates models.StudentUpdate) (models.Student, error) {
	var student models.Student

	//create query
	query := fmt.Sprintf(`SELECT * FROM public.update_student($1, $2, $3, $4, $5, $6, $7)`)

	row := r.db.QueryRow(query, username, studentUpdates.FirstName, studentUpdates.LastName, studentUpdates.PhoneNumber, studentUpdates.Email, studentUpdates.LastLogin, studentUpdates.MajorID)

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.LastLogin, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}
