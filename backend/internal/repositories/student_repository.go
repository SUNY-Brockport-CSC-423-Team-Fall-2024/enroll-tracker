package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
	"fmt"
)

type StudentRepository interface {
	CreateStudent(firstName string, lastName string, authId int, phoneNumber string, email string) (models.Student, error)
}

type PostgresStudentRepository struct {
	db *sql.DB
}

func (r *PostgresStudentRepository) CreateStudent(firstName string, lastName string, authId int, phoneNumber string, email string) (models.Student, error) {
	//create student
	var student models.Student

	//create query
	query := fmt.Sprintf(`SELECT first_name, last_name, auth_id, phone_number, email, FROM public.create_student($1, $2, $3, $4, $5)`)

	//execute function, instantiate student, close connection
	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email).Scan(&student.ID, &student.First_Name, &student.Last_Name, &student.Auth_ID, &student.Major_ID, &student.Phone_Number, &student.Email, &student.Last_Login, &student.Created_At, &student.Updated_At)

	return student, err
}
