package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type EnrollmentsRepository interface {
	EnrollStudent(courseID int, studentID int) (bool, error)
	UnenrollStudent(courseID int, studentID int) (bool, error)
	GetStudentsCourses(studentID int, isEnrolled *bool) ([]models.StudentsCourse, error)
	GetCoursesStudents(courseID int, isEnrolled *bool) ([]models.CoursesStudent, error)
	GetEnrollmentNumbers(courseID int, studentID int) (int, int, int, int, error)
}

type PostgresEnrollmentsRepository struct {
	db *sql.DB
}

func CreateEnrollmentsRepository(db *sql.DB) *PostgresEnrollmentsRepository {
	return &PostgresEnrollmentsRepository{db: db}
}

func (r *PostgresEnrollmentsRepository) EnrollStudent(courseID int, studentID int) (bool, error) {
	query := `SELECT * FROM public.enroll_student($1,$2)`

	if _, err := r.db.Exec(query, courseID, studentID); err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostgresEnrollmentsRepository) UnenrollStudent(courseID int, studentID int) (bool, error) {
	query := `SELECT * FROM public.unenroll_student($1,$2)`

	if _, err := r.db.Exec(query, courseID, studentID); err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostgresEnrollmentsRepository) GetStudentsCourses(studentID int, isEnrolled *bool) ([]models.StudentsCourse, error) {
	courses := make([]models.StudentsCourse, 0)

	query := `SELECT * FROM public.get_students_courses($1,$2)`

	rows, err := r.db.Query(query, studentID, isEnrolled)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		course := models.StudentsCourse{}
		if err := rows.Scan(&course.CourseID, &course.CourseName, &course.CourseDescription, &course.TeacherID, &course.MaxEnrollment, &course.NumCredits, &course.Status, &course.LastUpdated, &course.CreatedAt, &course.IsEnrolled, &course.UnenrolledDate, &course.EnrolledDate); err != nil {
			return courses, err
		}

		courses = append(courses, course)
	}
	return courses, nil
}

func (r *PostgresEnrollmentsRepository) GetCoursesStudents(courseID int, isEnrolled *bool) ([]models.CoursesStudent, error) {
	students := make([]models.CoursesStudent, 0)

	query := `SELECT * FROM public.get_courses_students($1,$2)`

	rows, err := r.db.Query(query, courseID, isEnrolled)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.CoursesStudent
		if err := rows.Scan(&student.StudentUsername, &student.StudentID, &student.FirstName, &student.LastName, &student.AuthID, &student.MajorID, &student.PhoneNumber, &student.Email, &student.CreatedAt, &student.UpdatedAt, &student.IsEnrolled, &student.EnrolledDate, &student.UnenrolledDate); err != nil {
			return students, err
		}

		students = append(students, student)
	}
	return students, nil
}

func (r *PostgresEnrollmentsRepository) GetEnrollmentNumbers(courseID int, studentID int) (int, int, int, int, error) {
	var numOfEnrolledStudents int
	var courseMaxEnrollment int
	var studentEnrollment int
	var numOfCourseCredits int

	query := `SELECT * FROM public.get_enrollment_numbers($1,$2)`

	err := r.db.QueryRow(query, courseID, studentID).Scan(&numOfEnrolledStudents, &courseMaxEnrollment, &studentEnrollment, &numOfCourseCredits)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return numOfEnrolledStudents, courseMaxEnrollment, studentEnrollment, numOfCourseCredits, nil
}
