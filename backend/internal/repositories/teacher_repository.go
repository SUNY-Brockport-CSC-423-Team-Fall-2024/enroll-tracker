package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type TeacherRepository interface {
	CreateTeacher(firstName string, lastName string, authId int, phoneNumber string, email string, office string) (models.Teacher, error)
	GetTeacher(username string) (models.Teacher, error)
	UpdateTeacher(username string, teacherUpdates models.TeacherUpdate) (models.Teacher, error)
}

type PostgresTeacherRepository struct {
	db *sql.DB
}

func CreatePostgresTeacherRepository(db *sql.DB) *PostgresTeacherRepository {
	return &PostgresTeacherRepository{db: db}
}

func (r *PostgresTeacherRepository) CreateTeacher(firstName string, lastName string, authId int, phoneNumber string, email string, office string) (models.Teacher, error) {
	var teacher models.Teacher

	query := `SELECT * FROM public.create_teacher($1,$2,$3,$4,$5,$6)`

	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email, office).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office)

	return teacher, err
}

func (r *PostgresTeacherRepository) GetTeacher(username string) (models.Teacher, error) {
	//create teacher
	var teacher models.Teacher

	//create query
	query := `SELECT * FROM public.get_teacher($1)`

	//execute function
	row := r.db.QueryRow(query, username)

	err := row.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, err
}

func (r *PostgresTeacherRepository) UpdateTeacher(username string, teacherUpdates models.TeacherUpdate) (models.Teacher, error) {
	var teacher models.Teacher

	//create query
	query := `SELECT * FROM public.update_teacher($1,$2,$3,$4,$5,$6)`

	row := r.db.QueryRow(query, username, teacherUpdates.FirstName, teacherUpdates.LastName, teacherUpdates.PhoneNumber, teacherUpdates.Email, teacherUpdates.Office)

	err := row.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, err
}
