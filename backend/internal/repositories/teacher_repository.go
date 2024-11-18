package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type TeacherRepository interface {
	CreateTeacher(firstName string, lastName string, authId int, phoneNumber string, email string, office string) (models.Teacher, error)
	GetTeachers(queryParams models.TeacherQueryParams) ([]models.Teacher, error)
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

	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email, office).Scan(&teacher.Username, &teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, err
}

func (r *PostgresTeacherRepository) GetTeachers(queryParams models.TeacherQueryParams) ([]models.Teacher, error) {
	//teacher array
	var teachers = make([]models.Teacher, 0)

	query := `SELECT * FROM public.get_teachers($1,$2,$3,$4,$5,$6,$7,$8)`

	rows, err := r.db.Query(query, queryParams.Limit, queryParams.Offset, queryParams.FirstName, queryParams.LastName, queryParams.Username, queryParams.Email, queryParams.PhoneNumber, queryParams.Office)
	if err != nil {
		return nil, err
	}
	//Be sure to close connection
	defer rows.Close()

	//Loop through returned rows
	for rows.Next() {
		teacher := models.Teacher{}
		if err := rows.Scan(&teacher.Username, &teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt); err != nil {
			return teachers, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (r *PostgresTeacherRepository) GetTeacher(username string) (models.Teacher, error) {
	//create teacher
	var teacher models.Teacher

	//create query
	query := `SELECT * FROM public.get_teacher($1)`

	//execute function
	row := r.db.QueryRow(query, username)

	err := row.Scan(&teacher.Username, &teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, err
}

func (r *PostgresTeacherRepository) UpdateTeacher(username string, teacherUpdates models.TeacherUpdate) (models.Teacher, error) {
	var teacher models.Teacher

	//create query
	query := `SELECT * FROM public.update_teacher($1,$2,$3,$4,$5,$6)`

	row := r.db.QueryRow(query, username, teacherUpdates.FirstName, teacherUpdates.LastName, teacherUpdates.PhoneNumber, teacherUpdates.Email, teacherUpdates.Office)

	err := row.Scan(&teacher.Username, &teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.AuthID, &teacher.PhoneNumber, &teacher.Email, &teacher.Office, &teacher.CreatedAt, &teacher.UpdatedAt)

	return teacher, err
}
