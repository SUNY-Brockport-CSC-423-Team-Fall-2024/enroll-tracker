package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"errors"
)

type CourseService struct {
	repository            repositories.CourseRepository
	courseMajorRepository repositories.CourseMajorRepository
}

func CreateCourseService(repo repositories.CourseRepository, cmRepo repositories.CourseMajorRepository) *CourseService {
	return &CourseService{repository: repo, courseMajorRepository: cmRepo}
}

func (s *CourseService) CreateCourse(courseCreation models.CourseCreation) (models.Course, error) {
	course, err := s.repository.CreateCourse(courseCreation)
	if err != nil {
		return course, err
	}
	return course, nil
}
func (s *CourseService) GetCourses(queryParams models.CourseQueryParams) ([]models.Course, error) {
	courses, err := s.repository.GetCourses(queryParams)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (s *CourseService) GetCourse(courseID int) (models.Course, error) {
	course, err := s.repository.GetCourse(courseID)
	if err != nil {
		return course, err
	}
	return course, nil
}

func (s *CourseService) UpdateCourse(courseID int, courseUpdates models.CourseUpdate) (bool, error) {
	//Validate updates
	if courseUpdates.MaxEnrollment != nil && (*courseUpdates.MaxEnrollment < 0 || *courseUpdates.MaxEnrollment > 100) {
		return false, errors.New("Max enrollment needs to be in the range 0-99")
	}
	if courseUpdates.NumCredits != nil && (*courseUpdates.NumCredits <= 0 || *courseUpdates.NumCredits > 6) {
		return false, errors.New("Num credits needs to be in the range 1-6")
	}

	//Set max enrollment to 0 if we are making a course inactive
	if courseUpdates.Status != nil && (*courseUpdates.Status == "inactive") {
		if courseUpdates.MaxEnrollment == nil {
			courseUpdates.MaxEnrollment = new(int)
		}
		*courseUpdates.MaxEnrollment = 0
	}

	success, err := s.repository.UpdateCourse(courseID, courseUpdates)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *CourseService) DeleteCourse(courseID int) (bool, error) {
	success, err := s.repository.DeleteCourse(courseID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}

func (s *CourseService) GetMajorsAssociatedWithCourse(courseID int, queryParams models.MajorQueryParams) ([]models.Major, error) {
	majors, err := s.courseMajorRepository.GetMajorsAssoicatedWithCourse(courseID, queryParams)
	if err != nil {
		return []models.Major{}, err
	}
	return majors, nil
}
