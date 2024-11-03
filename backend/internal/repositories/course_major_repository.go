package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type CourseMajorRepository interface {
	AddCourseToMajor(majorID int, courseID int) (bool, error)
	GetCoursesAssoicatedWithMajor(majorID int, queryParams models.CourseQueryParams) ([]models.Course, error)
	RemoveCourseFromMajor(majorID int, courseID int) (bool, error)
}

type PostgresCourseMajorRepository struct {
	db *sql.DB
}

func CreateCourseMajorRepository(db *sql.DB) *PostgresCourseMajorRepository {
	return &PostgresCourseMajorRepository{db: db}
}

func (r *PostgresCourseMajorRepository) AddCourseToMajor(majorID int, courseID int) (bool, error) {
	query := `SELECT * FROM public.add_course_to_major($1,$2)`

	if _, err := r.db.Exec(query, majorID, courseID); err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostgresCourseMajorRepository) GetCoursesAssoicatedWithMajor(majorID int, queryParams models.CourseQueryParams) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	query := `SELECT * FROM public.get_courses_associated_with_major($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	rows, err := r.db.Query(query, majorID, queryParams.Limit, queryParams.Offset, queryParams.Name, queryParams.Description, queryParams.TeacherID, queryParams.MaxEnrollment, queryParams.MinEnrollment, queryParams.MaxNumCredits, queryParams.MinNumCredits, queryParams.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.TeacherID, &course.MaxEnrollment, &course.NumCredits, &course.Status, &course.LastUpdated, &course.CreatedAt); err != nil {
			return courses, err
		}

		courses = append(courses, course)
	}
	return courses, nil
}

func (r *PostgresCourseMajorRepository) RemoveCourseFromMajor(majorID int, courseID int) (bool, error) {
	query := `SELECT * FROM public.remove_course_from_major($1,$2)`

	if _, err := r.db.Exec(query, majorID, courseID); err != nil {
		return false, err
	}
	return true, nil
}
