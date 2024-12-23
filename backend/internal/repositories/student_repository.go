package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"

	"github.com/lib/pq"
	_ "golang.org/x/tools/go/analysis/passes/ifaceassert"
)

type StudentRepository interface {
	CreateStudent(firstName string, lastName string, authId int, phoneNumber string, email string) (models.Student, error)
	GetStudents(queryParams models.StudentQueryParams) ([]models.Student, error)
	GetStudent(username string) (models.Student, error)
	UpdateStudent(username string, studentUpdates models.StudentUpdate) (models.Student, error)
	AddStudentToMajor(majorID int, studentID int) (bool, error)
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
	query := `SELECT * FROM public.create_student($1, $2, $3, $4, $5)`

	//execute function, instantiate student, close connection
	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email).Scan(&student.Username, &student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}

func (r *PostgresStudentRepository) GetStudents(queryParams models.StudentQueryParams) ([]models.Student, error) {
	//student array
	var students = make([]models.Student, 0)

	query := `SELECT * FROM public.get_students($1,$2,$3,$4,$5,$6,$7,$8)`

	rows, err := r.db.Query(query, queryParams.Limit, queryParams.Offset, queryParams.FirstName, queryParams.LastName, queryParams.Username, pq.Array(queryParams.Majors), queryParams.Email, queryParams.PhoneNumber)
	if err != nil {
		return nil, err
	}
	//Be sure to close connection
	defer rows.Close()

	//Loop through returned rows
	for rows.Next() {
		student := models.Student{}
		if err := rows.Scan(&student.Username, &student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt); err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *PostgresStudentRepository) GetStudent(username string) (models.Student, error) {
	//create student
	var student models.Student

	//create query
	query := `SELECT * FROM public.get_student($1)`

	//execute function
	row := r.db.QueryRow(query, username)

	err := row.Scan(&student.Username, &student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}

func (r *PostgresStudentRepository) UpdateStudent(username string, studentUpdates models.StudentUpdate) (models.Student, error) {
	var student models.Student

	//create query
	query := `SELECT * FROM public.update_student($1, $2, $3, $4, $5)`

	row := r.db.QueryRow(query, username, studentUpdates.FirstName, studentUpdates.LastName, studentUpdates.PhoneNumber, studentUpdates.Email)

	err := row.Scan(&student.Username, &student.ID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt)

	return student, err
}

func (r *PostgresStudentRepository) AddStudentToMajor(majorID int, studentID int) (bool, error) {
	var success bool

	query := `SELECT * FROM public.add_student_to_major($1,$2)`

	err := r.db.QueryRow(query, majorID, studentID).Scan(&success)
	if err != nil {
		return false, err
	}
	return success, nil
}
