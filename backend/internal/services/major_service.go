package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
)

type MajorService struct {
	repository            repositories.MajorRepository
	courseMajorRepository repositories.CourseMajorRepository
}

func CreateMajorService(repo repositories.MajorRepository, courseMajorRepo repositories.CourseMajorRepository) *MajorService {
	return &MajorService{repository: repo, courseMajorRepository: courseMajorRepo}
}

func (s *MajorService) CreateMajor(majorCreation models.MajorCreation) (models.Major, error) {
	major, err := s.repository.CreateMajor(majorCreation)
	if err != nil {
		return major, err
	}
	return major, nil
}
func (s *MajorService) GetMajors(queryParams models.MajorQueryParams) ([]models.Major, error) {
	majors, err := s.repository.GetMajors(queryParams)
	if err != nil {
		return majors, err
	}
	return majors, nil
}

func (s *MajorService) GetMajor(majorID int) (models.Major, error) {
	major, err := s.repository.GetMajor(majorID)
	if err != nil {
		return major, err
	}
	return major, nil
}

func (s *MajorService) UpdateMajor(majorID int, majorUpdates models.MajorUpdate) (bool, error) {
	success, err := s.repository.UpdateMajor(majorID, majorUpdates)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *MajorService) DeleteMajor(majorID int) (bool, error) {
	success, err := s.repository.DeleteMajor(majorID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}

func (s *MajorService) AddCourseToMajor(majorIDs []int, courseID int) (bool, error) {
	success, err := s.courseMajorRepository.AddCourseToMajor(majorIDs, courseID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}

func (s *MajorService) GetCoursesAssociatedWithMajor(majorID int, queryParams models.CourseQueryParams) ([]models.Course, error) {
	courses, err := s.courseMajorRepository.GetCoursesAssoicatedWithMajor(majorID, queryParams)
	if err != nil {
		return []models.Course{}, err
	}
	return courses, nil
}

func (s *MajorService) RemoveCourseFromMajor(majorID int, courseID int) (bool, error) {
	success, err := s.courseMajorRepository.RemoveCourseFromMajor(majorID, courseID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}
