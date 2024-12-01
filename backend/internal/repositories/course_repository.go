package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type CourseRepository interface {
	CreateCourse(courseCreation models.CourseCreation) (models.Course, error)
	GetCourses(queryParams models.CourseQueryParams) ([]models.Course, error)
	GetCourse(courseID int) (models.Course, error)
	UpdateCourse(courseID int, courseUpdates models.CourseUpdate) (bool, error)
	DeleteCourse(courseID int) (bool, error)
}

type PostgresCourseRepository struct {
	db *sql.DB
}

func CreateCourseRepository(db *sql.DB) *PostgresCourseRepository {
	return &PostgresCourseRepository{db: db}
}

func (r *PostgresCourseRepository) CreateCourse(courseCreation models.CourseCreation) (models.Course, error) {
	var course models.Course

	query := `SELECT * FROM public.create_course($1,$2,$3,$4,$5)`

	err := r.db.QueryRow(query, courseCreation.Name, courseCreation.Description, courseCreation.TeacherID, courseCreation.MaxEnrollment, courseCreation.NumCredits).Scan(&course.ID, &course.Name, &course.Description, &course.TeacherID, &course.MaxEnrollment, &course.NumCredits, &course.Status, &course.LastUpdated, &course.CreatedAt)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}
func (r *PostgresCourseRepository) GetCourses(queryParams models.CourseQueryParams) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	query := `SELECT * FROM public.get_courses($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	rows, err := r.db.Query(query, queryParams.Limit, queryParams.Offset, queryParams.Name, queryParams.Description, queryParams.TeacherID, queryParams.MaxEnrollment, queryParams.MinEnrollment, queryParams.MaxNumCredits, queryParams.MinNumCredits, queryParams.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.TeacherID, &course.CurrentEnrollment, &course.MaxEnrollment, &course.NumCredits, &course.Status, &course.LastUpdated, &course.CreatedAt); err != nil {
			return courses, err
		}

		courses = append(courses, course)
	}
	return courses, nil
}

func (r *PostgresCourseRepository) GetCourse(courseID int) (models.Course, error) {
	var course models.Course

	query := `SELECT * FROM public.get_course($1)`

	err := r.db.QueryRow(query, courseID).Scan(&course.ID, &course.Name, &course.Description, &course.TeacherID, &course.CurrentEnrollment, &course.MaxEnrollment, &course.NumCredits, &course.Status, &course.LastUpdated, &course.CreatedAt)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (r *PostgresCourseRepository) UpdateCourse(courseID int, courseUpdates models.CourseUpdate) (bool, error) {
	var success bool

	query := `SELECT * FROM public.update_course($1,$2,$3,$4,$5,$6)`

	err := r.db.QueryRow(query, courseID, courseUpdates.Description, courseUpdates.TeacherID, courseUpdates.MaxEnrollment, courseUpdates.NumCredits, courseUpdates.Status).Scan(&success)
	if err != nil {
		return false, err
	}
	if !success {
		return false, models.NoAffectedRows
	}
	return true, nil
}

func (r *PostgresCourseRepository) DeleteCourse(courseID int) (bool, error) {
	var success bool

	query := `SELECT * FROM public.delete_course($1)`

	err := r.db.QueryRow(query, courseID).Scan(&success)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}
