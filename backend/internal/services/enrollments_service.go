package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"errors"
)

type EnrollmentsService struct {
	repository repositories.EnrollmentsRepository
}

func CreateEnrollmentsService(repo repositories.EnrollmentsRepository) *EnrollmentsService {
	return &EnrollmentsService{repository: repo}
}

func (s *EnrollmentsService) EnrollStudent(courseID int, studentID int) (bool, error) {
	//Check current enrollment in class
	curEnrollment, maxEnrollment, studentEnrolledCredits, numOfCourseCredits, err := s.repository.GetEnrollmentNumbers(courseID, studentID)
	if err != nil {
		return false, err
	}

	if curEnrollment >= maxEnrollment {
		return false, errors.New("Course is full. Max enrollment reached.")
	}

	//Check for student's current num of enrolled credits
	if studentEnrolledCredits+numOfCourseCredits > 18 {
		return false, errors.New("Unable to enroll in course. Would put student over max number of course credits they can be enrolled in.")
	}

	success, err := s.repository.EnrollStudent(courseID, studentID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}

func (s *EnrollmentsService) UnenrollStudent(courseID int, studentID int) (bool, error) {
	success, err := s.repository.UnenrollStudent(courseID, studentID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}

func (s *EnrollmentsService) GetStudentsCourses(studentID int, isEnrolled *bool) ([]models.StudentsCourse, error) {
	courses, err := s.repository.GetStudentsCourses(studentID, isEnrolled)
	if err != nil {
		return []models.StudentsCourse{}, err
	}
	return courses, nil
}

func (s *EnrollmentsService) GetCoursesStudents(courseID int, isEnrolled *bool) ([]models.CoursesStudent, error) {
	students, err := s.repository.GetCoursesStudents(courseID, isEnrolled)
	if err != nil {
		return []models.CoursesStudent{}, err
	}
	return students, nil
}
